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
	panic("implement me")
}
