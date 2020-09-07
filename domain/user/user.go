package user

import "github.com/williamchang80/sea-apd/domain"

type User struct {
	domain.Base
	Username string `json:"username"`
}
