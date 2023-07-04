package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Issuer       string
	Audience     string
	Secret       string
	TokeExp      time.Duration
	RefreshExp   time.Duration
	CookieDomain string
	CookiePath   string
	CookieName   string
}

type jwtUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (a *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.Id)
	claims["aud"] = a.Audience
	claims["iss"] = a.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	//set jwt exp
	claims["exp"] = time.Now().UTC().Add(a.TokeExp).Unix()

	//create signed token
	signedAccessToken, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	//create refresh toke and set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.Id)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()

	//set refresh token exp
	refreshTokenClaims["exp"] = time.Now().UTC().Add(a.RefreshExp).Unix()

	//create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(a.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	//create token pairs and populate with signed tokens
	var tokenPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	//return token pairs ad nil error
	return tokenPairs, nil
}

func (a *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     a.CookieName,
		Path:     a.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(a.RefreshExp),
		MaxAge:   int(a.RefreshExp.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   a.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func (a *Auth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     a.CookieName,
		Path:     a.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   a.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func (a *Auth) GetTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
	w.Header().Add("Vary", "Authorization")

	authHeader := r.Header.Get("Authorization")

	// sanity check
	if authHeader == "" {
		return "", nil, errors.New("No auth headder")
	}

	// split the header on spaces, we should have 2 items
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", nil, errors.New("No auth headder")
	}

	// check for bearer key
	if headerParts[0] != "Bearer" {
		return "", nil, errors.New("No auth headder")
	}

	token := headerParts[1]

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.Secret), nil
	})

	if err != nil {

		if strings.HasPrefix(err.Error(), "token is expired by") {
			return "", nil, errors.New("expired token")
		}

		return "", nil, err
	}

	// make sure we issued the token
	if claims.Issuer != a.Issuer {
		return "", nil, errors.New("invalid issuer")
	}

	return token, claims, nil

}
