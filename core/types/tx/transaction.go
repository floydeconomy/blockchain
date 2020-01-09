package tx

// Transaction contains one transaction and it's information.
type Transaction struct {
	header *Header
	body   *Body
}

// Transactions an array of transaction
type Transactions []*Transaction

// Header returns the block header.
func (tx *Transaction) Header() *Header {
	return tx.header
}

// Body returns the block body.
func (tx *Transaction) Body() *Body {
	return tx.body
}
