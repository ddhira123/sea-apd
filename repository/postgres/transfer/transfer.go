package transfer

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/transfer"
)

type TransferRepository struct {
	db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) transfer.TransferRepository {
	return &TransferRepository{db: db}
}

func (t TransferRepository) GetTransferHistory(userId string) ([]transfer.Transfer, error) {
	var transfers []transfer.Transfer
	err := t.db.Where("user_id = ?", userId).Find(&transfers).Error
	if err != nil {
		return nil, err
	}
	return transfers,nil
}

func (t TransferRepository) CreateTransferHistory(transfer transfer.Transfer) error {
	if err := t.db.Create(&transfer).Error; err != nil {
		return err
	}
	return nil
}
