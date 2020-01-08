package tx

import (
  "github.com/floydeconomy/blockchain/types"
)

// Body contains the body information for a transaction
type Body struct {
  Clauses      []*Clause
}

// Clause contains the transaction information
type Clause struct {
  Payee   types.Address
  Payer   types.Address
  Amount  uint32
}
