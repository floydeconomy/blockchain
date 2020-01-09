package tx

// Header contains the header information for a transaction
type Header struct {
	ChainTag   byte
	BlockRef   uint64
	Expiration uint32
	Nonce      uint64
	Signature  []byte
}
