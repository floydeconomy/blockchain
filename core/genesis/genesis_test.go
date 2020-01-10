// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package genesis_test

import (
	"fmt"
	"testing"

	. "github.com/floydeconomy/blockchain/core/genesis"
	"github.com/stretchr/testify/assert"
)

func TestTestnetGenesis(t *testing.T) {
	// genesis
	genesis := NewTestnet()
	blk := genesis.BuildGenesisBlock()

	// block
	genesisParentID := blk.ParentID()
	genesisID := blk.ID()
	genesisNumber := blk.Number()

	// prints
	fmt.Println("ParentID: ", genesisParentID.String())
	fmt.Println("Genesis Block ID: ", genesisID.String())
	fmt.Println("Genesis Block Number: ", genesisNumber)

	// assets
	assert.Equal(t, NETWORK, genesis.Name())
	assert.Equal(t, GenesisBlockParentHash[:4], genesisParentID[:4])
	assert.Equal(t, genesisParentID, GenesisBlockParentHash)
	assert.Equal(t, uint32(0x0), genesisNumber)
}
