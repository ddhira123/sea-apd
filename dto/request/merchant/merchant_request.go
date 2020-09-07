package merchant

type UpdateMerchantBalanceRequest struct {
	Amount int `json:"amount"`
	MerchantId string `json:"merchant_id"`
}
