import { useState, useEffect} from 'react'
import { useParams, Link } from 'react-router-dom'

const Movie = () =>{
    const [movie, setMovie] = useState({})
    let {id} = useParams()


    useEffect(()=>{
        const  requestOptions = {
            method:"GET",
            headers:{
                'Content-Type': 'application/json',
            }
        }
 
        fetch(`/movie/${id}`, requestOptions)
        .then((resp) => resp.json())
        .then((data) => {
            if(data.error){
                console.log(data.error, 'cunt')
            }else{
                setMovie(data.movie)
            }
        })
        .catch((error)=>{
            console.log(error)
        })
    },[id])

    movie.genres = movie.genres ? Object.values(movie.genres) : [] 

    return(
        <>
        <div className="text-center">
            <h2>Movie: {movie.title}</h2>
            <small><em>{movie.release_date}, {movie.runtime} minutes, Rated {movie.mpaa_rating}</em></small>
            <br/>
            {
                movie.genres.map((g) => (
                    <span key={g.genre} className='badge bg-secondary me-2'>{g.genre}</span>
                ))
            }
            <hr/>
            {
                movie.image && 
                <div className='mb-3'>
                    <img src={`https://image.tmdb.org/t/p/w200/${movie.image}`}/>
                </div>
            }
            <p>{movie.description}</p>
            <Link to="/movies">Back</Link>
        </div>
        </>
    )
}

export default Movie