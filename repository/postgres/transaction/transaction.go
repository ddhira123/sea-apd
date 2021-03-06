package transaction

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/domain/transaction"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transaction.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t TransactionRepository) CreateCart(tr transaction.Transaction) error {
	err := t.db.Create(&tr).Error
	return err
}

func (t TransactionRepository) UpdateTransactionStatus(status string, id string) (*transaction.Transaction, error) {
	var tran transaction.Transaction
	err := t.db.Where("id = ?", id).Find(&tran).Update("status", status).Error
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
	err := t.db.Where("status IN (?)", requiredStatus).Where(
		"customer_id = ?", userId).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t TransactionRepository) GetMerchantRequestItem(merchantId string) ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	onRequestMerchantStatus := transaction_status.ToString(transaction_status.WAITING_DELIVERY)
	err := t.db.Model(&transactions).Where("status = ?", onRequestMerchantStatus).
		Where("merchant_id = ?", merchantId).
		Preload("ProductDetails").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t TransactionRepository) UpdateTransaction(tr transaction.Transaction) error {
	if err := t.db.Debug().Model(&tr).Where("id = ?", tr.ID).
		Updates(transaction.Transaction{
			BankNumber: tr.BankNumber,
			BankName:   tr.BankName,
			Amount:     tr.Amount,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (t TransactionRepository) AddCartItem(cart transaction.ProductTransaction) error {
	if err := t.db.Create(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (t TransactionRepository) RemoveCartItem(cart transaction.ProductTransaction) error {
	if err := t.db.Where("transactionId = ? AND productId = ?", cart.TransactionId, cart.ProductId).Delete(&transaction.ProductTransaction{}).Error; err != nil {
		return err
	}
	return nil
}

func (t TransactionRepository) UpdateCartItem(cart transaction.ProductTransaction) error {
	if err := t.db.Model(&cart).Where("transactionId = ? AND productId = ?", cart.TransactionId, cart.ProductId).Update(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (t TransactionRepository) GetCartItems(id string) ([]transaction.ProductTransaction, error) {
	var cartItems []transaction.ProductTransaction
	err := t.db.Where("transaction_id = ?", id).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
