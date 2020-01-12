package block

import (
	"bytes"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/rlp"
)

// Raw allows to partially decode components of a block.
type Raw []byte

// RawBlock struct of a raw block
type RawBlock struct {
	Raw   Raw
	cache struct {
		Header atomic.Value
		Body   atomic.Value
		Block  atomic.Value
	}
}

// DecodeHeader decode only the header.
func (r Raw) DecodeHeader() (*Header, error) {
	content, _, err := rlp.SplitList(r)
	if err != nil {
		return nil, err
	}

	var header Header
	if err := rlp.Decode(bytes.NewReader(content), &header); err != nil {
		return nil, err
	}
	return &header, nil
}

// DecodeBody decode only the body.
func (r Raw) DecodeBody() (*Body, error) {
	content, _, err := rlp.SplitList(r)
	if err != nil {
		return nil, err
	}

	var body Body
	if err := rlp.Decode(bytes.NewReader(content), &body); err != nil {
		return nil, err
	}
	return &body, nil
}

// NewRawBlock creates a new raw block
func NewRawBlock(raw Raw, block *Block) *RawBlock {
	rb := &RawBlock{Raw: raw}
	rb.cache.Header.Store(block.Header())
	rb.cache.Body.Store(block.Body())
	rb.cache.Block.Store(block)
	return rb
}

// Header parses the header
func (rb *RawBlock) Header() (*Header, error) {
	if cached := rb.cache.Header.Load(); cached != nil {
		return cached.(*Header), nil
	}

	h, err := rb.Raw.DecodeHeader()
	if err != nil {
		return nil, err
	}
	rb.cache.Header.Store(h)
	return h, nil
}

// Body parses the body
func (rb *RawBlock) Body() (*Body, error) {
	if cached := rb.cache.Body.Load(); cached != nil {
		return cached.(*Body), nil
	}
	b, err := rb.Raw.DecodeBody()
	if err != nil {
		return nil, err
	}
	rb.cache.Body.Store(b)
	return b, nil
}

// Block parses the block
func (rb *RawBlock) Block() (*Block, error) {
	if cached := rb.cache.Block.Load(); cached != nil {
		return cached.(*Block), nil
	}

	h, err := rb.Header()
	if err != nil {
		return nil, err
	}
	b, err := rb.Body()
	if err != nil {
		return nil, err
	}

	block := Compose(h, b.Txs)

	rb.cache.Block.Store(block)
	return block, nil
}
