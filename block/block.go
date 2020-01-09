package block

// Block contains almost all information about a block
// It's immutable.
type Block struct {
	header *Header
	body   *Body
}

// Header returns the block header.
func (b *Block) Header() *Header {
	return b.header
}

// Body returns the block body.
func (b *Block) Body() *Body {
	return b.body
}
