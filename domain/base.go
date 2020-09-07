package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

func (Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	scope.SetColumn("id", uuid)
	return nil
}

type Base struct {
	ID        string     `gorm:"primary_key;" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
