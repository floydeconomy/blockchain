package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/floydeconomy/blockchain/core/types/block"
	"github.com/vechain/thor/kv"
)

var (
	bestBlockPrefix = []byte("best") // ("best", block id) -> block
	blockPrefix     = []byte("b")    // ("b", block id) -> block
)

func saveBestBlock(w kv.Putter, id common.Hash, raw []byte) error {
	// BlockID
	if err := saveBestBlockID(w, id); err != nil {
		return err
	}

	// BlockRaw
	if err := saveBlockRaw(w, id, raw); err != nil {
		return err
	}

	return nil
}

// saveBestBlockID save the best block ID on trunk.
func saveBestBlockID(w kv.Putter, id common.Hash) error {
	return w.Put(bestBlockPrefix, id[:])
}

// loadBestBlockID returns the best block ID on trunk.
func loadBestBlockID(r kv.Getter) (common.Hash, error) {
	data, err := r.Get(bestBlockPrefix)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(data), nil
}

// saveBlockRaw save rlp encoded block raw data.
func saveBlockRaw(w kv.Putter, id common.Hash, raw block.Raw) error {
	return w.Put(append(blockPrefix, id[:]...), raw)
}

// loadBlockRaw load rlp encoded block raw data.
func loadBlockRaw(r kv.Getter, id common.Hash) (block.Raw, error) {
	return r.Get(append(blockPrefix, id[:]...))
}
