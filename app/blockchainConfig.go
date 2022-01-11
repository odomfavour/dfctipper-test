package app

type BlockchainConfig struct {
	BSCNode            string `long:"mainnet" env:"MAINNET_NODE_ADDRESS"`
	MasterAddressKey   string `long:"MASTER_ADDRESS_KEY" env:"MASTER_ADDRESS_KEY"`
	MasterAddress      string `long:"MASTER_ADDRESS" env:"MASTER_ADDRESS"`
}
