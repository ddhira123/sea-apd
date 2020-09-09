package role

func GetUserRoles() map[string]string {
	return map[string]string{
		"CUSTOMER": "customer",
		"MERCHANT": "merchant",
		"ADMIN":    "admin",
	}
}
