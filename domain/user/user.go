package user

import (
	"time"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
