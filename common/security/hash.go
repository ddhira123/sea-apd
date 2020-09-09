package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(p)
}

func IsMatchedPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
