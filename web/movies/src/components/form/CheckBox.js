import {forwardRef} from "react"

const CheckBox = forwardRef((props, ref) => {
    const {name, title, className, onChange, autoComplete, value, checked} = props
    return (
        <div className="mb-3">
            <input
                type='checkbox'
                className={className}
                id={name}
                ref={ref}
                name={name}
                onChange={onChange}
                checked={checked}
                autoComplete={autoComplete}
                value={value}
            />
            <label htmlFor={name} className="form-label">
                {title}
            </label>
        </div>
    )
})

export default CheckBox;