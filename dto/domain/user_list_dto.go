package domain

import "github.com/williamchang80/sea-apd/domain/user"

type UserListDto struct {
	Users []user.User `json:"users"`
}
