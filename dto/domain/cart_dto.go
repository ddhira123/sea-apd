package domain

import "github.com/williamchang80/sea-apd/domain/transaction"

type CartDto struct {
	CartItems []transaction.ProductTransaction `json:"cart_items"`
}
