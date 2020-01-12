package genesis

import "github.com/floydeconomy/blockchain/core/types/block"

const (
	testnet   string = "testnet"
	customnet string = "custom"
)

// Genesis specifies the header fields, state of a genesis block.
type Genesis struct {
	builder *Builder
	name    string
}

// BuildGenesisBlock build the genesis block.
// TODO: implement ID check like in thor
func (g *Genesis) BuildGenesisBlock() *block.Block {
	return g.builder.Build()
}

// Name returns network name.
func (g *Genesis) Name() string {
	return g.name
}
