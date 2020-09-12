package transaction

import (
	"errors"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/common/observer"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/domain/product"
	"github.com/williamchang80/sea-apd/domain/transaction"
	transaction2 "github.com/williamchang80/sea-apd/dto/request/transaction"
	"github.com/williamchang80/sea-apd/dto/request/transaction/converter"
)

var obs *TransactionObserver

type TransactionUsecase struct {
	tr              transaction.TransactionRepository
	merchantUseCase merchant.MerchantUsecase
	productUseCase  product.ProductUsecase
}

type TransactionObserver struct {
	observer.TransactionObservable
}

func NewTransactionUsecase(repo transaction.TransactionRepository,
	merchantUseCase merchant.MerchantUsecase, productUsecase product.
ProductUsecase) transaction.TransactionUsecase {
	obs = CreateObserverable()
	obs.attachObservers()
	return &TransactionUsecase{tr: repo,
		merchantUseCase: merchantUseCase,
		productUseCase:  productUsecase}
}

func convertTransactionRequestToDomain(t transaction2.CreateCartRequest) transaction.Transaction {
	return transaction.Transaction{
		Status:     transaction_status.ToString(transaction_status.ON_CARTS),
		CustomerId: t.UserId,
		MerchantId: t.MerchantId,
	}
}

func (i *TransactionObserver) attachObservers() {
	i.TransactionObservable.AddObserver(&UpdateMerchantBalanceObserver{})
	i.TransactionObservable.AddObserver(&NotifyAdminObserver{})
}

func CreateObserverable() *TransactionObserver {
	return &TransactionObserver{}
}

func (t TransactionUsecase) UpdateTransactionStatus(request transaction2.
UpdateTransactionRequest) error {
	status := transaction_status.ToString(request.Status)
	tran, err := t.tr.UpdateTransactionStatus(status, request.TransactionId)
	if err != nil {
		return err
	}
	if err := obs.NotifyAll(*tran, t.merchantUseCase); err != nil {
		return err
	}
	return nil
}

func (t TransactionUsecase) GetTransactionById(id string) (*transaction.Transaction, error) {
	tr, err := t.tr.GetTransactionById(id)
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (t TransactionUsecase) GetTransactionHistory(userId string) ([]transaction.
Transaction, error) {
	requiredStatusForTransactionHistory := transaction_status.GetStatusListForTransactionHistory()
	tr, err := t.tr.GetTransactionByRequiredStatus(requiredStatusForTransactionHistory, userId)
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (t TransactionUsecase) GetMerchantRequestItem(merchantId string) ([]transaction.Transaction, error) {
	tr, err := t.tr.GetMerchantRequestItem(merchantId)
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (t TransactionUsecase) PayTransaction(request transaction2.PaymentRequest) error {
	tr, err := t.tr.GetTransactionById(request.TransactionId)
	if err != nil {
		return err
	}
	transactionTotal, err := t.productUseCase.GetProductPriceTotal(*tr)
	if err != nil {
		return err
	}
	mergedTransaction := converter.MergePaymentRequestAndTransactionTotal(
		request, *tr, transactionTotal)
	if err := t.tr.UpdateTransaction(mergedTransaction); err != nil {
		return err
	}
	updateRequest := transaction2.UpdateTransactionRequest{
		TransactionId: mergedTransaction.ID,
		Status:        transaction_status.WAITING_CONFIRMATION,
	}
	if err := t.UpdateTransactionStatus(updateRequest); err != nil {
		return err
	}
	return nil
}

func ConvertCartRequest(req transaction2.CartRequest) transaction.ProductTransaction {
	return transaction.ProductTransaction{
		ProductId:     req.ProductId,
		TransactionId: req.TransactionId,
		Quantity:      req.Quantity,
	}
}

func CheckOnCart(t TransactionUsecase, transactionId string) error {
	trs, err := t.tr.GetTransactionById(transactionId)
	if err != nil {
		return err
	}
	if transaction_status.ParseToEnum(trs.Status) != transaction_status.ON_CARTS {
		return errors.New("Transaction's status is not on cart.")
	}
	return nil
}

func (t TransactionUsecase) AddCartItem(request transaction2.CartRequest) error {
	errr := CheckOnCart(t, request.TransactionId)
	if errr != nil {
		return errr
	}
	ci := ConvertCartRequest(request)
	err := t.tr.AddCartItem(ci)
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionUsecase) RemoveCartItem(request transaction2.CartRequest) error {
	errr := CheckOnCart(t, request.TransactionId)
	if errr != nil {
		return errr
	}
	ci := ConvertCartRequest(request)
	err := t.tr.RemoveCartItem(ci)
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionUsecase) UpdateCartItem(request transaction2.CartRequest) error {
	errr := CheckOnCart(t, request.TransactionId)
	if errr != nil {
		return errr
	}
	ci := ConvertCartRequest(request)
	err := t.tr.UpdateCartItem(ci)
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionUsecase) GetCartItems(id string) ([]transaction.ProductTransaction, error) {
	ci, err := t.tr.GetCartItems(id)
	if err != nil {
		return nil, err
	}
	return ci, nil
}

func (t TransactionUsecase) CreateCart(request transaction2.CreateCartRequest) error {
	tran := convertTransactionRequestToDomain(request)
	err := t.tr.CreateCart(tran)
	return err
}