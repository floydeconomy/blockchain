package tx

// Transaction contains one transaction and it's information.
type Transaction struct {
	header *Header
	body   *Body
}

// Transactions an array of transaction
type Transactions []*Transaction
