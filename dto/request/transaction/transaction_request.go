package transaction

type TransactionRequest struct {
	BankNumber string
	BankName   string
	Amount     int
	UserId 	   string
}

type UpdateTransactionRequest struct {
	Id string
	Status string
}
