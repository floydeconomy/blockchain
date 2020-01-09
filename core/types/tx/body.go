package tx

import (
	"github.com/vechain/thor/thor"
)

// Body contains the body information for a transaction
type Body struct {
	Clauses []*Clause
}

// Clause contains the transaction information
type Clause struct {
	Payee  thor.Address
	Payer  thor.Address
	Amount uint32
}
