package tx

// Header contains the header information for a transaction
type Header struct {
	body HeaderBody
}

// HeaderBody contains block headers.
type HeaderBody struct {
	Nonce uint64
}

// Nonce returns nonce value.
func (h *Header) Nonce() uint64 {
	return h.body.Nonce
}
