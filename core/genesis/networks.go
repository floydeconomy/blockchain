package genesis

// NewTestnet create mainnet genesis.
func NewTestnet() *Genesis {
	var (
		chainLaunchTime uint64 = uint64(1530014400) // 'Tue Jun 26 2018 20:00:00 GMT+0800 (CST)'
		extra           [28]byte
	)

	builder := new(Builder).
		Timestamp(chainLaunchTime).
		ExtraData(extra)

	return &Genesis{
		builder,
		testnet,
	}
}

// NewCustomnet create mainnet genesis.
func NewCustomnet(chainLaunchTime uint64, extra [28]byte) *Genesis {
	builder := new(Builder).
		Timestamp(chainLaunchTime).
		ExtraData(extra)

	return &Genesis{
		builder,
		customnet,
	}
}
