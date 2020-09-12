package auth

import "github.com/williamchang80/sea-apd/dto/response/base"

type LoginResponse struct {
	base.BaseResponse
	Token string `json:"token"`
}
