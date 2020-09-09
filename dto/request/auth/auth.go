package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	Name                 string `json:"name"`
	PasswordConfirmation string `json:"password_confirmation"`
}
