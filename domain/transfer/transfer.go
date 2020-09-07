package transfer

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/transfer"
)

type Transfer struct {
	domain.Base
	Amount     int    `json:"amount"`
	BankName   string `json:"bank_name"`
	BankNumber string `json:"bank_number"`
	UserId     string `json:"user_id"`
}

type TransferController interface {
	GetTransferHistory(ctx echo.Context) error
	CreateTransferHistory(ctx echo.Context) error
}

type TransferUsecase interface {
	GetTransferHistory(merchantId string) ([]Transfer, error)
	CreateTransferHistory(request transfer.CreateTransferHistoryRequest) error
}

type TransferRepository interface {
	GetTransferHistory(merchantId string) ([]Transfer, error)
	CreateTransferHistory(Transfer) error
}
