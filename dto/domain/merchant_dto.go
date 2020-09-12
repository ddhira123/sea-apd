package domain

import "github.com/williamchang80/sea-apd/domain/merchant"

type MerchantDto struct {
	Merchant *merchant.Merchant `json:"merchant"`
}
