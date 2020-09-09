package transaction

type TransactionRequest struct {
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
	Amount     int    `json:"amount"`
	UserId     string `json:"user_id"`
}

type UpdateTransactionRequest struct {
	TransactionId string `json:"transaction_id"`
	Status        string `json:"status"`
}
