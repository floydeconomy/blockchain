package tx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// Transaction contains one transaction and it's information.
type Transaction struct {
	header *Header
	body   *Body
}

// Transactions an array of transaction
type Transactions []*Transaction

// Header contains the header information for a transaction
type Header struct {
	Nonce uint64
}

// Body contains the body information for a transaction
// TODO: support more than one clause
type Body struct {
	Clause []*Clause
}

// Clause contains information of a cluase
type Clause struct {
	Recipient *common.Address
	Amount    *big.Int
	Data      []byte
}

// Header returns the block header.
func (tx *Transaction) Header() *Header {
	return tx.header
}

// Body returns the block body.
func (tx *Transaction) Body() *Body {
	return tx.body
}

// Nonce returns nonce value.
func (tx *Transaction) Nonce() uint64 {
	return tx.header.Nonce
}

// Clause returns a clause
func (tx *Transaction) Clause() []*Clause {
	return tx.body.Clause
}

// ID returns the transaction hash of the header, which is simply the keccak256 hash of its RLP Encoding
// TODO: implement more complex architecture of having RLP Hash seperate from ID
func (tx *Transaction) ID() (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, tx)
	hw.Sum(h[:0])
	return h
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
