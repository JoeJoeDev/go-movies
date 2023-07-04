import {forwardRef} from "react"

const TextArea = forwardRef((props, ref) => {
    const {name, title, type, className, placeholder, onChange, autoComplete, value, errorMsg, rows} = props
    return (
        <div className="mb-3">
            <label htmlFor={name} className="form-label">
                {title}
            </label>
            <textarea
                type={type}
                className={className}
                id={name}
                ref={ref}
                name={name}
                placeholder={placeholder}
                onChange={onChange}
                autoComplete={autoComplete}
                value={value}
                rows={rows}
            />
            <div className={props.errorDiv}>{errorMsg}</div>
        </div>
    )
})

export default TextArea;