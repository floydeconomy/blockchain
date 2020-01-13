package block

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/theblockchainbook/helpers/kv"
)

var (
	bestBlockPrefix = []byte("best") // ("best", block id) -> block
	blockPrefix     = []byte("b")    // ("b", block id) -> block
)

// SaveBestBlock save the best block ID on trunk.
func SaveBestBlock(w kv.Putter, id common.Hash, raw []byte) error {
	// BlockID
	if err := SaveBestBlockID(w, id); err != nil {
		return err
	}

	// BlockRaw
	if err := SaveBlockRaw(w, id, raw); err != nil {
		return err
	}

	return nil
}

// SaveBestBlockID save the best block ID on trunk.
func SaveBestBlockID(w kv.Putter, id common.Hash) error {
	return w.Put(bestBlockPrefix, id[:])
}

// LoadBestBlockID returns the best block ID on trunk.
func LoadBestBlockID(r kv.Getter) (common.Hash, error) {
	data, err := r.Get(bestBlockPrefix)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(data), nil
}

// SaveBlockRaw save rlp encoded block raw data.
func SaveBlockRaw(w kv.Putter, id common.Hash, raw Raw) error {
	return w.Put(append(blockPrefix, id[:]...), raw)
}

// LoadBlockRaw load rlp encoded block raw data.
func LoadBlockRaw(r kv.Getter, id common.Hash) (Raw, error) {
	return r.Get(append(blockPrefix, id[:]...))
}
