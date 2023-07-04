import React, {useEffect, useState} from 'react'
import { useNavigate, useOutletContext, useParams } from 'react-router-dom'
import Input from './form/Input'
import Select from './form/Select'
import TextArea from './form/TextArea'
import CheckBox from './form/CheckBox'
import Swal from 'sweetalert2'

function EditMovie() {

    const navigate = useNavigate()

    const {token} = useOutletContext()

    const [genres, setGenres] = useState([])
    const [error, setError] = useState(null)
    const [errors, setErrors] = useState([])

    const hasError =(key) =>{
        return errors.indexOf(key) !== -1
    }

    const [movie, setMovie] = useState(null)

    const mpaaRatings = [
        { id:"G", value:"G" },
        { id:"M", value:"M" },
        { id:"R", value:"R" },
        { id:"PG", value:"PG" },
        { id:"R18", value:"R18" },
        { id:"A", value:"A" }
]


    let {id} = useParams()
    

    useEffect(()=>{
        if(!token){
            navigate('/login')
            return
        }

        if(id){
            const  requestOptions = {
                method:"GET",
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            }
    
            fetch(`/movie/${id}`, requestOptions)
            .then((response) => {
                if(response.status === 401){
                    navigate('/login')
                }else{
                    return response
                }
            })
            .then((resp) => resp.json())
            .then((data) => {
                if(data.error){
                    console.log(data.error, 'cunt')
                }else{
                    setGenres(data.genres)
                    setMovie(m =>(data.movie))
                }
            })
            .catch((error)=>{
                console.log(error)
            })

        } else {

            const  requestOptions = {
                method:"GET",
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            }
    
            fetch(`/genres`, requestOptions)
            .then((response) => {
                if(response.status === 401){
                    navigate('/login')
                }else{
                    return response
                }
            })
            .then((resp) => resp.json())
            .then((data) => {
                if(data.error){
                    console.log(data.error)
                }else{
                    //setGenres(data)
                 if(data !== undefined){
                    const checks = []
                    
                    data.forEach((g)=>{
                        checks.push({id: g.id, checked: false, genre: g.genre })
                    })

                   
                 setMovie({
                        id: 0,
                        title: '',
                        release_date: '',
                        runtime: '',
                        mpaa_rating: '',
                        description: '',
                        genres: checks,
                        genres_array: []
                    })
                }

                }
            })
            .catch((error)=>{
                console.log(error)
            })

    }

    },[token, navigate, id])

    const handleSubmit= (e) =>{
        e.preventDefault()


        let errorsArray =[]

        let required = [
            {field: movie.title, name: "title"},
            {field: movie.release_date, name: "release_date"},
            {field: movie.runtime, name: "runtime"},
            {field: movie.description, name: "description"},
            {field: movie.mpaa_rating, name: "mpaa_rating"},
        ]

        required.forEach(function(obj){
            if(obj.field === ""){
                errorsArray.push(obj.name)
            }
        })

        if(movie.genres_array.length === 0){
            Swal.fire({
                title : "Error",
                text: "please select at least one genre",
                icon: "error",
                confirmButtonText:"Ok" 
            })
            errors.push("genres")
        }

        setErrors(errorsArray)

        debugger

        if(errors.length > 0){
            return false
        }

        
        // request handler

        // let payload = {
        //     email,
        //     password
        // }
  
        // const  requestOptions = {
        //     method:"POST",
        //     headers:{
        //         'Content-Type': 'application/json',
        //     },
        //     credentials: 'include',
        //     body: JSON.stringify(payload)
        // }
 
        // fetch('/authenticate', requestOptions)
        // .then((resp) => resp.json())
        // .then((data) => {
        //     if(data.error){
        //         setAlertClassName('alert-danger')
        //         setAlertMessage(data.message)
        //     }else{
        //         setToken(data.access_token)
        //         setAlertClassName('d-none')
        //         setAlertMessage('')
        //         navigate('/')
        //     }
        // })
        // .catch((error)=>{
        //     setAlertClassName('alert-danger')
        //     setAlertMessage(error)
        // })

    }

    const handleChange = () => (event) =>{
       let value = event.target.value
       let name = event.target.name

       setMovie({
        ...movie,
        [name]: value
       })

    }

    const handleChecked = (event, index) =>{

        let tmparr = movie.genres
        tmparr[index].checked = !tmparr[index].checked
        let tempIds = movie.genres_array
        if(!event.target.checked){
            tempIds.splice(tmparr.indexOf(event.target.value))
        }else{
            tempIds.push(parseInt(event.target.value, 10))
        }

        setMovie({
            ...movie,
            genres_array : tempIds
        })

    }

    return movie && (
        <div>
            <h1>Edit Movie</h1>
        
        <hr/>

        <pre>{JSON.stringify(movie, null, 3)}</pre>

        <form onSubmit={handleSubmit}>
            <input type="hidden" name="id" value={movie.id} id="id"></input>

        <Input
            title={"Title"}
            type={"text"}
            className={"form-control"}
            name={"title"}
            onChange={handleChange("title")}
            value={movie.title}
            errorDiv={hasError("title") ? "text-danger": "d-none"}
            errorMsg={"Please enter a title"}
        />
        <Input
            title={"Release Date"}
            type={"date"}
            className={"form-control"}
            name={"release_date"}
            onChange={handleChange("release_date")}
            value={movie.release_date}
            errorDiv={hasError("release_date") ? "text-danger": "d-none"}
            errorMsg={"Please enter a release date"}
        />
        <Input
            title={"Runtime"}
            type={"number"}
            className={"form-control"}
            name={"runtime"}
            onChange={handleChange("runtime")}
            value={movie.runtime}
            errorDiv={hasError("runtime") ? "text-danger": "d-none"}
            errorMsg={"Please enter a runtime"}
        />
        <Select
            title={"Rating"}
            className={"form-control"}
            name={"mpaa_rating"}
            options={mpaaRatings}
            onChange={handleChange("mpaa_rating")}
            value={movie.mpaa_rating}
            errorDiv={hasError("mpaa_rating") ? "text-danger": "d-none"}
            errorMsg={"Please choose a rating"}
            placeHolde={"Choose"}
        />
        <TextArea
            title={"Description"}
            className={"form-control"}
            name={"description"}
            onChange={handleChange("description")}
            rows={5}
            value={movie.description}
            errorDiv={hasError("description") ? "text-danger": "d-none"}
            errorMsg={"Please enter a description"}
        />

        <hr/>
        <h3>Genres</h3>
        {
            
                movie.genres && Array.isArray(movie.genres) && movie.genres.length > 0 && 
                <>
                {
                    
                    Array.from(movie.genres).map((g, index)=>  
                        <CheckBox
                            title={g.genre}
                            name={g.genre}
                            key={index}
                            id={"genre-" + index }
                            onChange={(e)=> handleChecked(e, index)}
                            value={g.id}
                            checked={movie.genres[index].checked}
                        />
                    )
                    
                }
                </>
         
        }

        <input
            type='submit'
            className='btn btn-primary'
            value='Save'
        />
    </form>
    </div>
    )
}

export default EditMovie