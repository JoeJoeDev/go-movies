package main

import "net/http"

func (app *application) enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Content-Type", "application/json")

		if request.Method == "OPTIONS" {
			writer.Header().Set("Access-Control-Allow-Credentials", "true")
			writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,PATCH,POST,DELETE,OPTIONS")
			writer.Header().Set("Access-Control-Allow-Headders", "Accept, Content-Type, X-CSRF-Token, Authoirzation, Accept")
			return
		} else {
			h.ServeHTTP(writer, request)
		}
	})
}


func (app * application) authRequired (next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_,_,err := app.auth.GetTokenFromHeaderAndVerify(w,r)
		if err != nil{
			w.WriteHeader(http.StatusUnauthorized)
		}
		next.ServeHTTP(w,r)

	})

}