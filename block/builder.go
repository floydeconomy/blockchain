package block

import (
  "github.com/floydeconomy/blockchain/tx"
  "github.com/vechain/thor/thor"
)

// Builder easy block building
type Builder struct {
	blockHeader  HeaderBody
	txs          tx.Transactions
}

// ParentID set parent id.
func (b *Builder) ParentID(id thor.Bytes32) *Builder {
	b.blockHeader.ParentID = id
	return b
}

// Timestamp set timestamp
func (b *Builder) Timestamp(ts uint64) *Builder {
  b.blockHeader.Timestamp = ts
  return b
}

// Beneficiary set beneficiary
func (b *Builder) Beneficiary(addr thor.Address) *Builder {
  b.blockHeader.Beneficiary = addr
  return b
}

// Transaction add a transaction.
func (b *Builder) Transaction(tx *tx.Transaction) *Builder {
	b.txs = append(b.txs, tx)
	return b
}

// Build build a block object.
func (b *Builder) Build() *Block {
  headerBody := HeaderBody{
    ParentID: b.blockHeader.ParentID,
    Beneficiary: b.blockHeader.Beneficiary,
    Timestamp: b.blockHeader.Timestamp,
  }
	header := Header{
    body: headerBody,
  }
  body := Body{Txs: b.txs}
	return &Block{
		header: &header,
		body:   &body,
	}
}
