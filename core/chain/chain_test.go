package chain_test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/floydeconomy/blockchain/core/chain"
	"github.com/floydeconomy/blockchain/core/genesis"
	"github.com/floydeconomy/blockchain/core/types/block"
	"github.com/floydeconomy/blockchain/core/types/tx"
	"github.com/stretchr/testify/assert"
	"github.com/theblockchainbook/helpers/lvldb"
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

func TestErrorGenesisBlock(t *testing.T) {
	kv, _ := lvldb.NewMem()

	// Test IsGenesisBlock
	blockNotGenesis := new(block.Builder).
		ParentID(common.BytesToHash([]byte("0"))).
		Build()
	_, err := chain.New(kv, blockNotGenesis)
	assert.NotNil(t, err)

	// Test HasTransaction
	var genesisBlockParentHash common.Hash = common.Hash{0xff, 0xff, 0xff, 0xff}
	tx := new(tx.Builder).
		Clause(tx.NewClause(nil)).
		Build()
	blockWithTransaction := new(block.Builder).
		Transaction(tx).
		ParentID(genesisBlockParentHash).
		Build()
	_, err = chain.New(kv, blockWithTransaction)
	assert.NotNil(t, err)

	// TODO: Test HasValidTimestamp
}

// For testing purposes, ensure deterministic block creation by not providing timestamp
func NewBlock(parent *block.Block) *block.Block {
	return new(block.Builder).ParentID(parent.ID()).Build()
}
