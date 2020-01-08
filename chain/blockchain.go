package chain

import (
	"github.com/floydeconomy/blockchain/block"
)

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*block.Block
}

// // AddBlock saves provided data as a block in the blockchain
// func (chain *Blockchain) AddBlock(data string) {
// 	builder = block.Builder{
// 		blockHeader: {
//
// 		},
// 	}
// 	builder = block.Builder()
//
// 	chain.Blocks = append(chain.Blocks, newBlock)
// }
//
// // NewBlockchain creates a new Blockchain with genesis Block
// func NewBlockchain() *Blockchain {
// 	return &Blockchain{[]*Block{NewGenesisBlock()}}
// }
