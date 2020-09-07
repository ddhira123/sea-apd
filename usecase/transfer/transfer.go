package transfer

import (
	"errors"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/domain/transfer"
	merchant2 "github.com/williamchang80/sea-apd/dto/request/merchant"
	request "github.com/williamchang80/sea-apd/dto/request/transfer"
)

type TransferUsecase struct {
	repo            transfer.TransferRepository
	merchantUsecase merchant.MerchantUsecase
}

func ConvertToDomain(request request.CreateTransferHistorysRequest) transfer.Transfer {
	return transfer.Transfer{
		Amount:     request.Amount,
		BankName:   request.BankName,
		BankNumber: request.BankNumber,
		MerchantId: request.MerchantId,
	}
}

func NewTransferUsecase(repo transfer.TransferRepository, usecase merchant.MerchantUsecase) transfer.TransferUsecase {
	return &TransferUsecase{repo: repo, merchantUsecase: usecase}
}

func (t TransferUsecase) GetTransferHistory(merchantId string) ([]transfer.Transfer, error) {
	transfers, err := t.repo.GetTransferHistory(merchantId)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

func validateMerchantBalanceAmount(amount int, balance int) error {
	if balance+amount < 0 {
		return errors.New("deposit amount cannot be more than wallet")
	}
	return nil
}

func (t TransferUsecase) CreateTransferHistory(request request.CreateTransferHistorysRequest) error {
	balance, err := t.merchantUsecase.GetMerchantBalance(request.MerchantId)
	if err != nil {
		return err
	}
	if err := validateMerchantBalanceAmount(request.Amount, balance); err != nil {
		return err
	}
	if err := t.repo.CreateTransferHistory(ConvertToDomain(request)); err != nil {
		return err
	}
	updateMerchantBalanceRequest := merchant2.UpdateMerchantBalanceRequest{
		Amount:     request.Amount,
		MerchantId: request.MerchantId,
	}
	if err := t.merchantUsecase.UpdateMerchantBalance(updateMerchantBalanceRequest); err != nil {
		return err
	}
	return nil
}
