package observer

import (
	"github.com/williamchang80/sea-apd/domain/transaction"
)

type TransactionObservable struct {
	Observers []TransactionObserver
}

func (o *TransactionObservable) AddObserver(obs TransactionObserver) {
	o.Observers = append(o.Observers, obs)
}

func (o *TransactionObservable) NotifyAll(transaction transaction.Transaction) error {
	for _, ob := range o.Observers {
		if err := ob.Update(transaction);err != nil {
			return err
		}
	}
	return nil
}
