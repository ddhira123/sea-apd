package user

import "github.com/williamchang80/sea-apd/common/constants/user_role"

type UpdateUserRoleRequest struct {
	Role   user_role.UserRole
	UserId string
}
