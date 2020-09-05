package transaction

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/transaction"
	"github.com/williamchang80/sea-apd/domain/user"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transaction.TransactionRepository {
	if db != nil {
		d := db.AutoMigrate(&transaction.Transaction{}, &user.User{})
		d.Model(&transaction.Transaction{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	}
	return &TransactionRepository{db: db}
}

func (t TransactionRepository) CreateTransaction(tr transaction.Transaction) error {
	err := t.db.Create(&tr).Error
	return err
}

func (t TransactionRepository) UpdateTransactionStatus(status string, id string) error {
	var tran transaction.Transaction
	err := t.db.First(&tran, id).Update("status", status).Error
	return err
}

func (t TransactionRepository) GetTransactionById(id string) (*transaction.Transaction, error) {
	var tran transaction.Transaction
	err := t.db.First(&tran, id).Error
	if err != nil {
		return nil, err
	}
	return &tran, nil
}

func (t TransactionRepository) GetTransactionByRequiredStatus(requiredStatus []string, userId string) ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	err := t.db.Where("status IN (?)", requiredStatus).Where("user_id = ?", userId).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
