package models

// SignUp struct to describe register a new user.
type SignUp struct {
	Email     string `json:"email" validate:"required,email,lte=255"`
	Firstname string `json:"firstname" validate:"required,lte=255"`
	Lastname  string `json:"lastname" validate:"required,lte=255"`
	About     string `json:"about" validate:"lte=255"`
	Password  string `json:"password" validate:"required,lte=255"`
	Role      string `json:"role" validate:"required,lte=25" default1:"user"`
}

// SignIn struct to describe login user.
type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
