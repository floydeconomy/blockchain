package tx

// Transaction the actual transaction itself
type Transaction struct {
	header *Header
	body   *Body
}
