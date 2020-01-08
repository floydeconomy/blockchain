package block

import (
  "encoding/binary"
  "sync/atomic"
  "golang.org/x/crypto/sha3"

  "github.com/vechain/thor/thor"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/rlp"
)

// Header holds header information.
type Header struct {
  body HeaderBody

  cache struct {
    id atomic.Value
  }
}

// HeaderBody contains block headers.
type HeaderBody struct {
  ParentID    thor.Bytes32
  Timestamp   uint64
  Beneficiary thor.Address
  Signature   []byte
}

// ParentID returns id of parent block.
func (h *Header) ParentID() thor.Bytes32 {
	return h.body.ParentID
}

// Number returns sequential number of this block.
func (h *Header) Number() uint32 {
	// inferred from parent id
	return Number(h.body.ParentID) + 1
}

// Number extract block number from block id.
func Number(blockID thor.Bytes32) uint32 {
	// first 4 bytes are over written by block number (big endian).
	return binary.BigEndian.Uint32(blockID[:])
}

// Timestamp returns timestamp of this block.
func (h *Header) Timestamp() uint64 {
	return h.body.Timestamp
}

// Beneficiary returns the beneficiary
func (h *Header) Beneficiary() thor.Address {
  return h.body.Beneficiary
}

// // ID computes id of block.
// // The block ID is defined as: blockNumber + hash(signingHash, signer)[4:].
// func (h *Header) ID() (id thor.Bytes32) {
// 	if cached := h.cache.id.Load(); cached != nil {
// 		return cached.(thor.Bytes32)
// 	}
// 	defer func() {
// 		// overwrite first 4 bytes of block hash to block number.
// 		binary.BigEndian.PutUint32(id[:], h.Number())
// 		h.cache.id.Store(id)
// 	}()
//
// 	signer, err := h.Signer()
// 	if err != nil {
// 		return
// 	}
//
// 	hw := thor.NewBlake2b()
// 	hw.Write(h.SigningHash().Bytes())
// 	hw.Write(signer.Bytes())
// 	hw.Sum(id[:0])
//
// 	return
// }


// ID returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
func (h *Header) ID() common.Hash {
	return Hash(h)
}

// Hash creating the hash from keccak256 library
func Hash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
