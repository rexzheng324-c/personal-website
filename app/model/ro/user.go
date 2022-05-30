package ro

type RegisterUserBody struct {
	// username
	// Required: true
	// Example: rex
	Username string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	// password
	// Required: true
	// Example: 123456
	Password string `json:"password" validate:"required,min=6,max=120" label:"密码"`
}

// swagger:parameters RegisterUser
type RegisterUser struct {
	// In: body
	Body RegisterUserBody
}

type LoginUserBody struct {
	// username
	// Required: true
	// Example: rex
	Username string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	// password
	// Required: true
	// Example: 123456
	Password string `json:"password" validate:"required,min=6,max=120" label:"密码"`
}

// swagger:parameters LoginUserBody
type LoginUser struct {
	// In: body
	Body LoginUserBody
}
