package block

import (
	"encoding/binary"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/floydeconomy/blockchain/core/types/tx"
	"golang.org/x/crypto/sha3"
)

// Block contains almost all information about a block
// It's immutable.
type Block struct {
	header *Header
	body   *Body
}

// Body contains the body information for a transaction
type Body struct {
	Txs tx.Transactions
}

// Header holds header information.
type Header struct {
	ParentID    common.Hash
	Timestamp   uint64
	Beneficiary common.Address
	Signature   []byte

	cache struct {
		id atomic.Value
	}
}

// Header returns the block header.
func (b *Block) Header() *Header {
	return b.header
}

// Body returns the block body.
func (b *Block) Body() *Body {
	return b.body
}

// ParentID returns id of parent block.
func (b *Block) ParentID() common.Hash {
	return b.header.ParentID
}

// Number returns sequential number of this block.
func (b *Block) Number() uint32 {
	return binary.BigEndian.Uint32(b.header.ParentID[:]) + 1
}

// Timestamp returns timestamp of this block.
func (b *Block) Timestamp() uint64 {
	return b.header.Timestamp
}

// Beneficiary returns the beneficiary
func (b *Block) Beneficiary() common.Address {
	return b.header.Beneficiary
}

// ID returns the block hash of the header, which is simply the keccak256 hash of its RLP Encoding
func (b *Block) ID() (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, b)
	hw.Sum(h[:0])
	return h
}

// Transactions returns a copy of transactions.
func (b *Block) Transactions() tx.Transactions {
	return append(tx.Transactions(nil), b.body.Txs...)
}
