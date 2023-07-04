import {forwardRef} from "react"

const Input = forwardRef((props, ref) => {
    const {name, title, type, className, placeholder, onChange, autoComplete, value, errorMsg} = props
    return (
        <div className="mb-3">
            <label htmlFor={name} className="form-label">
                {title}
            </label>
            <input
                type={type}
                className={className}
                id={name}
                ref={ref}
                name={name}
                placeholder={placeholder}
                onChange={onChange}
                autoComplete={autoComplete}
                value={value}
            />
            <div className={props.errorDiv}>{errorMsg}</div>
        </div>
    )
})

export default Input;