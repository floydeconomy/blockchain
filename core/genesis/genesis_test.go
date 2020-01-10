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
	GenesisBlockHash := common.Hash{0xff, 0xff, 0xff, 0xff} //so, genesis number is 0

	gene := NewTestnet()
	blk := gene.BuildGenesisBlock()
	id := blk.ParentID()

	fmt.Println("Parent ID: ", blk.ParentID())
	fmt.Println("Block ID: ", blk.ID())

	assert.Equal(t, "testnet", gene.Name())
	assert.Equal(t, GenesisBlockHash[:4], id[:4])
	assert.NotEqual(t, id, blk.ID())
}
