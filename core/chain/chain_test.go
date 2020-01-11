package chain_test

import (
	"fmt"
	"testing"

	"github.com/floydeconomy/blockchain/core/chain"
	"github.com/floydeconomy/blockchain/core/genesis"
	"github.com/floydeconomy/blockchain/core/types/block"
	"github.com/stretchr/testify/assert"
	"github.com/vechain/thor/lvldb"
)

func TestTestnetGenesis(t *testing.T) {
	kv, _ := lvldb.NewMem()
	g := genesis.NewTestnet()
	genesisBlock := g.BuildGenesisBlock()
	ch, err := chain.New(kv, genesisBlock)
	if err != nil {
		panic(err)
	}

	// Test Chain Genesis Block Initialization
	assert.Equal(t, genesisBlock, ch.BestBlock())
	assert.Equal(t, genesisBlock, ch.GenesisBlock())

	// Block Header
	id := genesisBlock.ID()
	fmt.Println("Genesis Block ID : ", id)
	rawBlockHeader, err := ch.GetBlockHeader(id)
	assert.Nil(t, err)
	assert.Equal(t, genesisBlock.Header().ParentID, rawBlockHeader.ParentID)

	// Set Blocks
	b0 := ch.GenesisBlock()
	b1 := NewBlock(b0)
	b2 := NewBlock(b1)
	b3 := NewBlock(b2)
	b4 := NewBlock(b3)

	tests := []struct {
		newBlock *block.Block
	}{
		{b1},
		{b2},
		{b3},
		{b4},
	}

	for _, tt := range tests {
		fmt.Println("New Block        : ", tt.newBlock.ID())
		err := ch.AddBlock(tt.newBlock)
		assert.Nil(t, err)
		assert.Equal(t, tt.newBlock.ID(), ch.BestBlock().ID())
		assert.Equal(t, b0.ID(), ch.GenesisBlock().ID())
	}
}

// For testing purposes, ensure deterministic block creation by not providing timestamp
func NewBlock(parent *block.Block) *block.Block {
	return new(block.Builder).ParentID(parent.ID()).Build()
}
