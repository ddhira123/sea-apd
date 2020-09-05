package payment

import (
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/domain/merchant"
)

type Payment struct {
	domain.Base
	Amount int `json:"amount"`
	BankName string `json:"bank_name"`
	BankNumber string `json:"bank_number"`
	Merchant merchant.Merchant `json:"merchant"`
	MerchantId string `json:"merchant_id"`
}
