package block

import (
	"encoding/binary"
	"io"
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

	cache struct {
		id atomic.Value
	}
}

// Body contains the body information for a block
type Body struct {
	Txs tx.Transactions
}

// Header holds header information.
type Header struct {
	ParentID    common.Hash
	Timestamp   uint64
	Beneficiary common.Address
	Signature   []byte
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
// TODO: implement block number
func (b *Block) ID() (h common.Hash) {
	if hash := b.cache.id.Load(); hash != nil {
		return hash.(common.Hash)
	}
	defer func() {
		// overwrite first 4 bytes of block hash to block number.
		binary.BigEndian.PutUint32(h[:], b.Number())
		b.cache.id.Store(h)
	}()

	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, b)
	hw.Sum(h[:0])
	return h
}

// Transactions returns a copy of transactions.
func (b *Block) Transactions() tx.Transactions {
	return append(tx.Transactions(nil), b.body.Txs...)
}

// Compose compose a block with all needed components
// Note: This method is usually to recover a block by its portions. To build up a block, use a Builder.
func Compose(header *Header, txs tx.Transactions) *Block {
	return &Block{
		header: header,
		body:   &Body{Txs: append(tx.Transactions(nil), txs...)},
	}
}

// EncodeRLP implements rlp.Encoder.
func (b *Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		b.header,
		b.body,
	})
}

// // DecodeRLP implements rlp.Decoder.
// func (b *Block) DecodeRLP(s *rlp.Stream) error {
// 	_, size, _ := s.Kind()
// 	payload := struct {
// 		header Header
// 		body   Body
// 	}{}
//
// 	if err := s.Decode(&payload); err != nil {
// 		return err
// 	}
//
// 	*b = Block{
// 		header: &payload.Header,
// 		txs:    payload.Txs,
// 	}
// 	b.cache.size.Store(metric.StorageSize(rlp.ListSize(size)))
// 	return nil
// }
