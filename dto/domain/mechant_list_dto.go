package domain

import "github.com/williamchang80/sea-apd/domain/merchant"

type MerchantListDto struct {
	Merchants []merchant.Merchant `json:"merchants"`
}
