// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package genesis_test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/floydeconomy/blockchain/core/genesis"
	"github.com/stretchr/testify/assert"
)

func TestTestnetGenesis(t *testing.T) {
	var (
		rootHash   common.Hash = common.Hash{0xff, 0xff, 0xff, 0xff}
		rootNumber uint32      = uint32(0x0)
		network    string      = "testnet"
	)

	// genesis
	genesis := NewTestnet()
	blk := genesis.BuildGenesisBlock()

	// block
	genesisParentID := blk.ParentID()
	genesisID := blk.ID()
	genesisNumber := blk.Number()

	// prints
	fmt.Println("ParentID: ", genesisParentID)
	fmt.Println("Genesis Block ID: ", genesisID)
	fmt.Println("Genesis Block Number: ", genesisNumber)

	// assets
	assert.Equal(t, network, genesis.Name())
	assert.Equal(t, rootHash[:4], genesisParentID[:4])
	assert.Equal(t, rootHash, genesisParentID)
	assert.Equal(t, rootNumber, genesisNumber)
}
