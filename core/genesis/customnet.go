package genesis

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
