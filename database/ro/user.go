package ro

type CreateUserBody struct {
	// username
	// Required: true
	// Example: rex
	Username string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	// username
	// Required: true
	// Example: 123456
	Password string `json:"password" validate:"required,min=6,max=120" label:"密码"`
}

// swagger:parameters CreateUser
type CreateUser struct {
	// In: body
	Body CreateUserBody
}
