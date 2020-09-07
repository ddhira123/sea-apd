package transfer

type CreateTransferHistoryRequest struct {
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
	Amount     int    `json:"amount"`
	MerchantId string `json:"merchant_id"`
}
