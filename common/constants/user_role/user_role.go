package user_role

type UserRole int

const (
	CUSTOMER = iota
	MERCHANT
	ADMIN
	OTHER
)

var UserRoleList = []string{
	"customer",
	"merchant",
	"admin",
}

func ToString(ts UserRole) string {
	if ts < CUSTOMER || ts > ADMIN {
		return ""
	}
	return UserRoleList[ts]
}

func ParseToEnum(src string) UserRole {
	userRoleMap := map[string]UserRole{
		"customer": CUSTOMER,
		"merchant": MERCHANT,
		"admin":    ADMIN,
		"other":    OTHER,
	}
	if val, exist := userRoleMap[src]; exist {
		return userRoleMap[string(rune(val))]
	}
	return userRoleMap["other"]
}

