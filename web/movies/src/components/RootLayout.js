import React, {useState, useEffect} from 'react'
import { Outlet, Link, useNavigate } from 'react-router-dom'
import Alert from './Alert'
export default function RootLayout() {

  const navigate = useNavigate()

  const [token, setToken] = useState("")
  const [alertMessage, setAlertMessage] = useState("")
  const [alertClassName, setAlertClassName] = useState("d-none")
  

  useEffect(() => {
    if(token === ""){
      const  requestOptions = {
        method:"GET",
        headers:{
            'Content-Type': 'application/json',
        },
        credentials: 'include',
      }

      fetch('/refresh', requestOptions)
      .then((resp) => resp.json())
      .then((data) => {
          if(data.accessToken){
            setToken(data.access_token)
          }else{
            console.log("No access token provided")
            setToken("")
          }
      })
      .catch((error)=>{
        console.log("Problem logging in", error)
      })
    }
  },[token])

  const logout = () => {
    const  requestOptions = {
      method:"GET",
      headers:{
          'Content-Type': 'application/json',
      },
      credentials: 'include',
    }

    fetch('/logout', requestOptions)
    .catch((error)=>{
      console.log("Problem logging out", error)
    })
    .finally(()=>{
      setToken("")
    })
    navigate("/login")
  }
  
  return (
    <div className="container">
    <div className="row">
      <div className="col">
        <h1 className="mt-3">Movies!</h1>
      </div>
      <div className="col text-end"> 
      { token === ""
        ? <Link to="/login"><span className="badge bg-success">Login</span></Link>
        : <Link onClick={logout}><span className="badge bg-danger">Logout</span></Link>
      }
      </div>
    </div>
    <div className="row">
      <div className="col-md-2">
        <nav>
          <div className="list-group">
            <Link to="/" className="list-group-item list-group-item-action">Home</Link>
            <Link to="/movies" className="list-group-item list-group-item-action">Movies</Link>
            <Link to="/genres" className="list-group-item list-group-item-action">Genres</Link>
            { token !== "" &&
              <>
                <Link to="/admin/movie" className="list-group-item list-group-item-action">Add Movie</Link>
                <Link to="/admin" className="list-group-item list-group-item-action">Manage Catalogue</Link>
              </>
            }
          </div>
        </nav>
      </div>
      <div className="col-md-10">
        <Alert
          message={alertMessage}
          className={alertClassName}
        />
        <Outlet context={{
          token, 
          setToken,
          alertClassName,
          setAlertClassName,
          alertMessage,
          setAlertMessage,
        }}/>
      </div>
      </div>

    </div>
  )
}
