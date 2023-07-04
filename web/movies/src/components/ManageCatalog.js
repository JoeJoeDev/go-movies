import { useState, useEffect} from 'react'
import { Link, useNavigate, useOutletContext} from 'react-router-dom'

const ManageCatalog = () =>{

    const [movies, setMovies] = useState([])

    const navigate = useNavigate()
    const {token} = useOutletContext()


    useEffect(()=>{

        if(!token){
            navigate('/login')
            return
        }

        const  requestOptions = {
            method:"GET",
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            }
        }
 
        fetch('/admin/movies', requestOptions)
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
                setMovies(data)
            }
        })
        .catch((error)=>{
            console.log(error)
        })
    },[token, navigate])

    return(
        <>
        <div className="text-center">
            <h2>Manage Catalog</h2>
            <hr/>

            <table className='table table-striped table-hover'>
                <thead>
                    <tr>
                        <th>Movie</th>
                        <th>Release Date</th>
                        <th>Rating</th>
                    </tr>
                </thead>
                <tbody>
                    {
                     movies.map((movie)=>(
                            <tr key={movie.id}>
                                <td> <Link to={`/admin/movies/${movie.id}`}>{movie.title}</Link></td>
                                <td>{movie.release_date}</td>
                                <td>{movie.mpaa_rating}</td>
                            </tr>
                     ))

                    }
                </tbody>
            </table>
        
        </div>
        </>
    )
}

export default ManageCatalog