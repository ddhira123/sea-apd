package admin

type AdminRequest struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
