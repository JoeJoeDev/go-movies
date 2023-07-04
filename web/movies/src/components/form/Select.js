import {forwardRef} from "react"

const Select = forwardRef((props, ref) => {
    const {name, title, type, className, placeholder, onChange, autoComplete, value, errorMsg, options} = props
    return (
        <div className="mb-3">
            <label htmlFor={name} className="form-label">
                {title}
            </label>
            <select
                className={className}
                id={name}
                ref={ref}
                name={name}
                onChange={onChange}
                autoComplete={autoComplete}
                value={value}
            >
            
            <option value="">{placeholder}</option>
            {
                options.map((option)=>{
                   return( <option key={option.id} value={option.id}>
                                {option.value}
                            </option>
                    )
                })
            }

            </select>
            <div className={props.errorDiv}>{errorMsg}</div>
        </div>
    )
})

export default Select;