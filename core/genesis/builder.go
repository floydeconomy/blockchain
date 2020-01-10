package genesis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/floydeconomy/blockchain/core/types/block"
)

// Builder helper to build genesis block.
type Builder struct {
	header ChainConfig
	body   BlockConfig
}

// ChainConfig defines the chain configuration
// TODO: implement this feature to handle hard fork switches over the chain configuration
type ChainConfig struct {
	Config *params.ChainConfig `json:"config"`
}

// BlockConfig contains all information for the block configuration
type BlockConfig struct {
	Nonce     uint64   `json:"nonce"`
	Timestamp uint64   `json:"timestamp"`
	ExtraData [28]byte `json:"extraData"`
}

// Timestamp set timestamp.
func (b *Builder) Timestamp(t uint64) *Builder {
	b.body.Timestamp = t
	return b
}

// ExtraData set extra data, which will be put into last 28 bytes of genesis parent id.
func (b *Builder) ExtraData(data [28]byte) *Builder {
	b.body.ExtraData = data
	return b
}

// Build build genesis block according to presets.
// TODO: implement stateCreator
func (b *Builder) Build() *block.Block {
	GenesisBlockHash := common.Hash{0xff, 0xff, 0xff, 0xff} //so, genesis number is 0
	copy(GenesisBlockHash[4:], b.body.ExtraData[:])
	return new(block.Builder).
		ParentID(GenesisBlockHash).
		Timestamp(b.body.Timestamp).
		Build()
}
