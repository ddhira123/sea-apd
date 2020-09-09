package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/williamchang80/sea-apd/domain/user"
	"os"
	"strings"
	"time"
)

func GenerateToken(user *user.User) (string, error) {
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID,
		"user_role":  user.Role,
		"exp":        time.Now().Add(time.Second).Unix(),
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

func GetValidBearerToken(token string) string {
	const bearerTokenPrefix = "Bearer "
	if strings.HasPrefix(token, bearerTokenPrefix) {
		return strings.TrimPrefix(token, bearerTokenPrefix)
	}
	return token
}

func IsValidTokenLifetime(t string) bool {
	secretKey := GetSecretKey()
	validBearerToken := GetValidBearerToken(t)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(validBearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if token == nil || err != nil || claims.Valid() != nil {
		return false
	}
	return true
}
