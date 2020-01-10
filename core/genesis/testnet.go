// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package genesis

// NETWORK constant depicting the network type
const NETWORK = "testnet"

// NewTestnet create mainnet genesis.
func NewTestnet() *Genesis {
	var (
		chainLaunchTime uint64 = uint64(1530014400) // 'Tue Jun 26 2018 20:00:00 GMT+0800 (CST)'
		extra           [28]byte
	)

	copy(extra[:], "Salute, Ethereum & Vechain")

	builder := new(Builder).
		Timestamp(chainLaunchTime).
		ExtraData(extra)

	return &Genesis{
		builder,
		NETWORK,
	}
}
