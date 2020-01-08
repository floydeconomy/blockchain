package main

import (
	"fmt"

  "github.com/floydeconomy/blockchain/chain"
)

func main() {
	bc := chain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.header.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.header.Hash)
		fmt.Println()
	}
}
