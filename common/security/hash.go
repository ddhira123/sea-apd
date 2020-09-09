package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/williamchang80/sea-apd/domain/user"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

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

func GenerateToken(user *user.User) (string, error) {
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID,
		"user_role":  user.Role,
		"expired_at": time.Now().Add(time.Hour).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secretKey := GetSecretKey()
	token, err := at.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetSecretKey() string {
	return os.Getenv("SECRET_AUTH_KEY")
}
