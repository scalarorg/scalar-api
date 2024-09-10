package types

type Staker struct {
	StakerAddress             string
	StakerPubkey              string
	ProtocolPubkey            string
	CovenantPubkeys           []string
	Qorum                     int
	Tag                       string
	Version                   int
	ChainID                   string
	ChainIdUserAddress        string
	ChainSmartContractAddress string
	MintingAmount             int
}
