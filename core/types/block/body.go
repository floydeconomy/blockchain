package block

import (
	"github.com/floydeconomy/blockchain/core/types/tx"
)

// Body contains the body information for a transaction
type Body struct {
	Txs tx.Transactions
}

// Transactions returns a copy of transactions.
func (b *Body) Transactions() tx.Transactions {
	return append(tx.Transactions(nil), b.Txs...)
}
