package transaction

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/transaction"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transaction.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t TransactionRepository) CreateTransaction(tr transaction.Transaction) error {
	err := t.db.Create(&tr).Error
	return err
}

func (t TransactionRepository) UpdateTransactionStatus(status string, id string) (*transaction.Transaction, error) {
	var tran transaction.Transaction
	err := t.db.First(&tran, id).Update("status", status).Error
	return &tran, err
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
