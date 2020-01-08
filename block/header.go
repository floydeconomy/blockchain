package block

import (
  "github.com/floydeconomy/blockchain/types"
  "encoding/binary"
)

// Header holds header information.
type Header struct {
  body HeaderBody
}

// HeaderBody contains block headers.
type HeaderBody struct {
  ParentID    types.Bytes32
  Timestamp   uint64
  Beneficiary types.Address
  Signature   []byte
}

// ParentID returns id of parent block.
func (h *Header) ParentID() types.Bytes32 {
	return h.body.ParentID
}

// Number returns sequential number of this block.
func (h *Header) Number() uint32 {
	// inferred from parent id
	return Number(h.body.ParentID) + 1
}

// Number extract block number from block id.
func Number(blockID types.Bytes32) uint32 {
	// first 4 bytes are over written by block number (big endian).
	return binary.BigEndian.Uint32(blockID[:])
}

// Timestamp returns timestamp of this block.
func (h *Header) Timestamp() uint64 {
	return h.body.Timestamp
}

// Beneficiary returns the beneficiary
func (h *Header) Beneficiary() types.Address {
  return h.body.Beneficiary
}
