package tx

// Builder to make it easy to build transaction.
type Builder struct {
	header Header
	body   Body
}

// Clause append a clause
func (b *Builder) Clause(c *Clause) *Builder {
	b.body.Clauses = append(b.body.Clauses, c)
	return b
}
