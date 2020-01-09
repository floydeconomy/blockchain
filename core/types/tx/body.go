package tx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Body contains the body information for a transaction
// TODO: support more than one clause
type Body struct {
	Clause *Clause
}

// Clause contains information of a cluase
type Clause struct {
	Recipient *common.Address
	Amount    *big.Int
	Data      []byte
}

// NewClause create a new clause instance.
func NewClause(to *common.Address) *Clause {
	if to != nil {
		// make a copy of 'to'
		cpy := *to
		to = &cpy
	}
	return &Clause{
		Recipient: to,
		Amount:    &big.Int{},
		Data:      nil,
	}
}
