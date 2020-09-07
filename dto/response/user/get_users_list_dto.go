package user

import (
	"github.com/williamchang80/sea-apd/dto/domain"
)

type GetUsersResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data domain.UserListDto `json:"data"`
}
