package user

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/user"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) CreateUser(user user.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetUserByEmail(email string) (*user.User, error) {
	var user user.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) UpdateUserRole(role string, userId string) error {
	err := u.db.Model(&user.User{}).Where("id = ?", userId).Updates(user.User{Role: role}).Error
	return err
}

func (u UserRepository) GetUserById(userId string) (*user.User, error) {
	var user user.User
	if err := u.db.Where("id = ?", userId).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) UpdateUser(user user.User) error {
	return u.db.Model(&user).Updates(&user).Error
}
