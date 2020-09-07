package domain

import "github.com/williamchang80/sea-apd/domain/user"

type UserDto struct {
	User user.User `json:"user"`
}
