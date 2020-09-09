package merchant

type UpdateMerchantBalanceRequest struct {
	Amount     int    `json:"amount"`
	MerchantId string `json:"merchant_id"`
}

type MerchantRequest struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	UserId  string `json:"user_id"`
	Brand   string `json:"brand"`
	Address string `json:"address"`
}
