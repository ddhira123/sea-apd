package observer

import (
	merchant3 "github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/domain/transaction"
)

type TransactionObservable struct {
	Observers []TransactionObserver
}

func (o *TransactionObservable) AddObserver(obs TransactionObserver) {
	o.Observers = append(o.Observers, obs)
}

func (o *TransactionObservable) NotifyAll(transaction transaction.Transaction,
	u merchant3.MerchantUsecase) error {
	for _, ob := range o.Observers {
		if err := ob.Update(transaction, u); err != nil {
			return err
		}
	}
	return nil
}
