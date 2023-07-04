import { useState, useEffect} from 'react'
import { Link } from 'react-router-dom'

const Movies = () =>{

    const [movies, setMovies] = useState([])

    useEffect(()=>{
      
        const  requestOptions = {
            method:"GET",
            headers:{
                'Content-Type': 'application/json',
            }
        }
 
        fetch('/movies', requestOptions)
        .then((resp) => resp.json())
        .then((data) => {
            if(data.error){
                console.log(data.error)
            }else{
                setMovies(data)
            }
        })
        .catch((error)=>{
            console.log(error)
        })
    },[])

    return(
        <>
        <div className="text-center">
            <h2>Movies</h2>
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
                                <td> <Link to={`/movie/${movie.id}`}>{movie.title}</Link></td>
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

export default Movies