package block

import (
	"github.com/floydeconomy/blockchain/core/types/tx"
)

// Body contains the body information for a transaction
type Body struct {
	Txs []*tx.Transaction
}

// Transactions returns transactions.
// TODO: use copy instead
func (b *Body) Transactions() []*tx.Transaction {
	return b.Txs
}
