package merchant

import "github.com/williamchang80/sea-apd/common/constants/merchant_status"

type UpdateMerchantBalanceRequest struct {
	Amount     int    `json:"amount"`
	MerchantId string `json:"merchant_id"`
}

type MerchantRequest struct {
	Name    string `json:"name"`
	UserId  string `json:"user_id"`
	Brand   string `json:"brand"`
	Address string `json:"address"`
}

type UpdateMerchantApprovalStatusRequest struct {
	Status     merchant_status.MerchantApprovalStatus `json:"status"`
	MerchantId string                                 `json:"merchant_id"`
}
