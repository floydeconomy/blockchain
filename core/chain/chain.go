package chain

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/floydeconomy/blockchain/core/types/block"
	"github.com/theblockchainbook/helpers/cache"
	"github.com/vechain/thor/co"
	"github.com/vechain/thor/kv"
)

var (
	blockCacheLimit      = 512
	errNotFound          = errors.New("not found")
	errBlockExist        = errors.New("block already exists")
	errTimestampNotFound = errors.New("timestamp missing")
)

// Chain describes a persistent block chain.
// It's thread-safe.
// TODO: implement reorg and forks
type Chain struct {
	kv kv.GetPutter

	genesisBlock *block.Block
	bestBlock    *block.Block
	tag          byte

	caches caches

	rw   sync.RWMutex
	tick co.Signal
}

type caches struct {
	blocks   *cache.Manager
	receipts *cache.Manager
}

// New create an instance of Chain.
func New(kv kv.GetPutter, genesisBlock *block.Block) (*Chain, error) {
	if !genesisBlock.IsGenesisBlock() {
		return nil, errors.New("genesis number != 0")
	}
	if genesisBlock.HasTransaction() {
		return nil, errors.New("genesis block should not have transactions")
	}
	if !genesisBlock.HasValidTimestamp() {
		return nil, errors.New("genesis block time cannot be in the future")
	}

	var bestBlock *block.Block
	if bestBlockID, err := block.LoadBestBlockID(kv); err != nil {
		if !kv.IsNotFound(err) {
			return nil, err
		}
		if bestBlock, err = handleEmptyGenesisBlock(kv, genesisBlock); err != nil {
			return nil, err
		}
	} else {
		if bestBlock, err = handleNotEmptyGenesisBlock(kv, bestBlockID); err != nil {
			return nil, err
		}
	}

	rawBlocksCache := cache.NewCache(blockCacheLimit, func(key interface{}) (interface{}, error) {
		raw, err := block.LoadBlockRaw(kv, key.(common.Hash))
		if err != nil {
			return nil, err
		}
		return &block.RawBlock{Raw: raw}, nil
	})

	return &Chain{
		kv:           kv,
		genesisBlock: genesisBlock,
		bestBlock:    bestBlock,
		tag:          genesisBlock.ID()[31],
		caches: caches{
			blocks: rawBlocksCache,
		},
	}, nil
}

func handleEmptyGenesisBlock(kv kv.GetPutter, genesisBlock *block.Block) (*block.Block, error) {
	// No Genesis Block has been found
	raw, err := rlp.EncodeToBytes(genesisBlock)
	if err != nil {
		return nil, err
	}

	// Create Batch
	batch := kv.NewBatch()

	// Save to kv
	if err := block.SaveBestBlock(batch, genesisBlock.ID(), raw); err != nil {
		return nil, err
	}

	if err := batch.Write(); err != nil {
		return nil, err
	}

	return genesisBlock, nil
}

func handleNotEmptyGenesisBlock(kv kv.GetPutter, bestBlockID common.Hash) (*block.Block, error) {
	// Load Block
	raw, err := block.LoadBlockRaw(kv, bestBlockID)
	if err != nil {
		return nil, err
	}
	bestBlock, err := (&block.RawBlock{Raw: raw}).Block()
	if err != nil {
		return nil, err
	}
	return bestBlock, nil
}

// GenesisBlock returns the genesis block of the chain
func (c *Chain) GenesisBlock() *block.Block {
	return c.genesisBlock
}

// BestBlock returns the genesis block of the chain
func (c *Chain) BestBlock() *block.Block {
	return c.bestBlock
}

// AddBlock add a new block into block chain.
func (c *Chain) AddBlock(newBlock *block.Block) error {
	if !newBlock.HasValidTimestamp() {
		return errors.New("genesis block time cannot be in the future")
	}

	c.rw.Lock()
	defer c.rw.Unlock()

	newBlockID := newBlock.ID()

	// Check if block exists
	if _, err := c.getBlockHeader(newBlockID); err != nil {
		if !c.IsNotFound(err) {
			return err
		}
	} else {
		// block already there
		return errBlockExist
	}

	// Check if parent block exists
	_, err := c.getBlockHeader(newBlock.ParentID())
	if err != nil {
		if c.IsNotFound(err) {
			return errors.New("parent missing")
		}
		return err
	}

	// Encode Block
	raw, err := rlp.EncodeToBytes(newBlock)
	if err != nil {
		return err
	}

	// Save to kv
	batch := c.kv.NewBatch()
	if err := block.SaveBestBlock(batch, newBlockID, raw); err != nil {
		return err
	}

	if err := batch.Write(); err != nil {
		return err
	}

	c.bestBlock = newBlock
	c.caches.blocks.Add(newBlockID, block.NewRawBlock(raw, newBlock))
	c.tick.Broadcast()

	return nil
}

// IsNotFound returns if an error means not found.
func (c *Chain) IsNotFound(err error) bool {
	return err == errNotFound || c.kv.IsNotFound(err)
}

// GetBlockHeader get block header by block id.
func (c *Chain) GetBlockHeader(id common.Hash) (*block.Header, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.getBlockHeader(id)
}

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
