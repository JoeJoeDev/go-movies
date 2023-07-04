import React, {useState} from 'react'
import {useNavigate, useOutletContext} from 'react-router-dom'
import Input from './form/Input'

export default function Login() {
    const [email, setEmail] = useState("admin@example.com")
    const [password, setPassword] = useState("secret")


    const {setToken} = useOutletContext()
    const {setAlertClassName} = useOutletContext()
    const {setAlertMessage} = useOutletContext()

    const navigate = useNavigate()

    const handleSubmit= (e) =>{
        e.preventDefault()
        
        // request handler

        let payload = {
            email,
            password
        }
  
        const  requestOptions = {
            method:"POST",
            headers:{
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify(payload)
        }
 
        fetch('/authenticate', requestOptions)
        .then((resp) => resp.json())
        .then((data) => {
            if(data.error){
                setAlertClassName('alert-danger')
                setAlertMessage(data.message)
            }else{
                setToken(data.access_token)
                setAlertClassName('d-none')
                setAlertMessage('')
                navigate('/')
            }
        })
        .catch((error)=>{
            setAlertClassName('alert-danger')
            setAlertMessage(error)
        })

    }
    return (
    <div className="col-md-6 offset-md-3">
        <h2>Login</h2>
        <hr/>

        <form onSubmit={handleSubmit}>
            <Input
                title="Email Address"
                type="email"
                className="form-control"
                name="email"
                autoComplete="email-new"
                onChange={((e)=>setEmail(e.target.value))}
                value="admin@example.com"
            />
            <Input
                title="Password"
                type="password"
                className="form-control"
                name="password"
                autoComplete="password-new"
                onChange={((e)=>setPassword(e.target.value))}
                value="secret"
            />

            <hr/>
            <input
                type='submit'
                className='btn btn-primary'
                value='Login'
            />
        </form>


    </div>
  )
}
