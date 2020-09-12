package user

import "github.com/williamchang80/sea-apd/common/constants/user_role"

type UpdateUserRoleRequest struct {
	Role   user_role.UserRole
	UserId string
}

type UpdateUserRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	NewEmail    string `json:"new_email"`
	OldEmail    string `json:"old_email"`
	UserId      string `json:"user_id"`
}
