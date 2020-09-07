package transfer

import "github.com/williamchang80/sea-apd/domain"

type Transfer struct {
	domain.Base
	Amount int `json:"amount"`
	BankName string `json:"bank_name"`
	BankNumber string `json:"bank_number"`
}

type TransferController interface {

}

type TransferUsecase interface {

}

type TransferRepository interface {
	GetTransferHistory(merchantId string)
}