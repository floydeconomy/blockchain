package tx

// Builder to make it easy to build transaction.
// TODO: clause should support array
type Builder struct {
	header Header
	clause *Clause
}

// Clause appends a clause
func (b *Builder) Clause(c *Clause) *Builder {
	b.clause = c
	return b
}

// Nonce set nonce
func (b *Builder) Nonce(nonce uint64) *Builder {
	b.header.body.Nonce = nonce
	return b
}

// Build build tx object
func (b *Builder) Build() *Transaction {
	body := Body{Clause: b.clause}
	tx := Transaction{
		header: &b.header,
		body:   &body,
	}
	return &tx
}
