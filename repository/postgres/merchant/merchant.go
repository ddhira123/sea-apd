package merchant

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/merchant"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) merchant.MerchantRepository {
	if db != nil {
		db.AutoMigrate(&merchant.Merchant{})
	}
	return &MerchantRepository{db: db}
}

func (m MerchantRepository) UpdateMerchantBalance(amount int, merchantId string) error {
	var merchant merchant.Merchant
	if err := m.db.Model(&merchant).Find(&merchant, merchantId).Update("balance", gorm.Expr("balance + ?",
		amount)).Error; err != nil {
		return err
	}
	return nil
}

func (m MerchantRepository) GetMerchantBalance(merchantId string) (int, error) {
	var merchant merchant.Merchant
	if err := m.db.Find(&merchant, merchantId).Error; err != nil {
		return 0, err
	}
	return merchant.Balance, nil
}
