package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/floydeconomy/blockchain/core/types/block"
)

func (c *Chain) getBlockHeader(id common.Hash) (*block.Header, error) {
	raw, err := c.getRawBlock(id)
	if err != nil {
		return nil, err
	}
	return raw.Header()
}

func (c *Chain) getRawBlock(id common.Hash) (*block.RawBlock, error) {
	raw, err := c.caches.blocks.GetOrLoad(id)
	if err != nil {
		return nil, err
	}
	return raw.(*block.RawBlock), nil
}

func (c *Chain) getBlockBody(id common.Hash) (*block.Body, error) {
	raw, err := c.getRawBlock(id)
	if err != nil {
		return nil, err
	}
	return raw.Body()
}

func (c *Chain) getBlock(id common.Hash) (*block.Block, error) {
	raw, err := c.getRawBlock(id)
	if err != nil {
		return nil, err
	}
	return raw.Block()
}
