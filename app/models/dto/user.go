package dto

type GetUserBody struct {
	// username
	// Required: true
	// Example: rex
	Username string `json:"username"`
	// password
	// Required: true
	// Example: 123456
	Password string `json:"password"`
	// role
	// Required: true
	// Example: 0
	Role int `json:"role"`
}

// swagger:response GetUser
type GetUser struct {
	// In: body
	Body GetUserBody
}
