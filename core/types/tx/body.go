package tx

import (
	"math/big"

	"github.com/vechain/thor/thor"
)

// Body contains the body information for a transaction
type Body struct {
	Clauses []*Clause
}

// Clause contains the transaction information
type Clause struct {
	To    *thor.Address `rlp:"nil"`
	Value *big.Int
	Data  []byte
}
