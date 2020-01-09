package block

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/floydeconomy/blockchain/core/types/tx"
)

// Builder easy block building
type Builder struct {
	header Header
	txs    []*tx.Transaction
}

// ParentID set parent id.
func (b *Builder) ParentID(id common.Hash) *Builder {
	b.header.ParentID = id
	return b
}

// Timestamp set timestamp
func (b *Builder) Timestamp(ts uint64) *Builder {
	b.header.Timestamp = ts
	return b
}

// Beneficiary set beneficiary
func (b *Builder) Beneficiary(addr common.Address) *Builder {
	b.header.Beneficiary = addr
	return b
}

// Transaction add a transaction.
func (b *Builder) Transaction(tx *tx.Transaction) *Builder {
	b.txs = append(b.txs, tx)
	return b
}

// Build build a block object.
func (b *Builder) Build() *Block {
	header := Header{
		ParentID:    b.header.ParentID,
		Beneficiary: b.header.Beneficiary,
		Timestamp:   b.header.Timestamp,
	}
	body := Body{
		Txs: b.txs,
	}
	return &Block{
		header: &header,
		body:   &body,
	}
}
