package merchant

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/merchant"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) merchant.MerchantRepository {
	return &MerchantRepository{db: db}
}

func (m MerchantRepository) UpdateMerchantBalance(amount int, merchantId string) error {
	var merchant merchant.Merchant
	if err := m.db.Model(&merchant).Where("id = ?", merchantId).Find(&merchant).Update("balance", gorm.Expr("balance + ?",
		amount)).Error; err != nil {
		return err
	}
	return nil
}

func (m MerchantRepository) GetMerchantBalance(merchantId string) (int, error) {
	var merchant merchant.Merchant
	if err := m.db.Where("id = ?", merchantId).Find(&merchant).Error; err != nil {
		return 0, err
	}
	return merchant.Balance, nil
}

func (m MerchantRepository) GetMerchants() ([]merchant.Merchant, error) {
	var merchants []merchant.Merchant
	err := m.db.Find(&merchants).Error
	if err != nil {
		return nil, err
	}
	return merchants, nil
}

func (m MerchantRepository) GetMerchantById(merchantId string) (*merchant.Merchant, error) {
	var merchant merchant.Merchant
	err := m.db.Where("id = ?", merchantId).Find(&merchant).Limit(1).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (m MerchantRepository) GetMerchantByUserId(userId string) (*merchant.Merchant, error) {
	var merchant merchant.Merchant
	err := m.db.Where("user_id = ?", userId).Find(&merchant).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (m MerchantRepository) RegisterMerchant(merchant merchant.Merchant) (*merchant.Merchant, error) {
	if err := m.db.Create(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (m MerchantRepository) UpdateMerchantApprovalStatus(merchantId string, status string) error {
	if err := m.db.Model(&merchant.Merchant{}).Where("id = ?", merchantId).Update(
		merchant.Merchant{Approval: status}).Error; err != nil {
		return err
	}
	return nil
}

func (m MerchantRepository) UpdateMerchant(merchantId string, merch merchant.Merchant) error {
	if err := m.db.Model(&merch).Where("id = ?",merchantId).
		Updates(&merch).Error; err != nil {
		return err
	}
	return nil
}