const Alert = ({className, message}) =>{

    console.log(className, message)
    return(
        <div className={"alert " + className} role="alert">
            {message}
        </div>
    )

}

export default Alert