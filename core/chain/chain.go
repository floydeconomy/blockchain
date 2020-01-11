package chain

import (
	"errors"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/floydeconomy/blockchain/core/types/block"
	"github.com/vechain/thor/co"
	"github.com/vechain/thor/kv"
)

// Chain describes a persistent block chain.
// It's thread-safe.
type Chain struct {
	kv kv.GetPutter

	genesisBlock *block.Block
	bestBlock    *block.Block
	tag          byte

	rw   sync.RWMutex
	tick co.Signal
}

// New create an instance of Chain.
func New(kv kv.GetPutter, genesisBlock *block.Block) (*Chain, error) {
	if genesisBlock.Number() != 0 {
		return nil, errors.New("genesis number != 0")
	}
	if len(genesisBlock.Transactions()) != 0 {
		return nil, errors.New("genesis block should not have transactions")
	}
	if genesisBlock.Timestamp() > uint64(time.Now().UnixNano()) {
		return nil, errors.New("genesis block time should be in the past")
	}

	var bestBlock *block.Block
	if bestBlockID, err := loadBestBlockID(kv); err != nil {
		if !kv.IsNotFound(err) {
			return nil, err
		}
		bestBlock, err = handleEmptyGenesisBlock(kv, genesisBlock)
		if err != nil {
			return nil, err
		}
	} else {
		bestBlock, err = handleNotEmptyGenesisBlock(kv, bestBlockID)
		if err != nil {
			return nil, err
		}
	}

	return &Chain{
		kv:           kv,
		genesisBlock: genesisBlock,
		bestBlock:    bestBlock,
		tag:          genesisBlock.ID()[31],
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
	err = saveBestBlock(batch, genesisBlock.ID(), raw)
	if err != nil {
		return nil, err
	}

	if err := batch.Write(); err != nil {
		return nil, err
	}

	return genesisBlock, nil
}

func handleNotEmptyGenesisBlock(kv kv.GetPutter, bestBlockID common.Hash) (*block.Block, error) {
	// Load Block
	raw, err := loadBlockRaw(kv, bestBlockID)
	if err != nil {
		return nil, err
	}
	bestBlock, err := (&rawBlock{raw: raw}).Block()
	if err != nil {
		return nil, err
	}
	return bestBlock, nil
}
