package app

type BlockchainConfig struct {
	// Blockchain
	BSCNode            string `long:"MAINNET_NODE_ADDRESS" env:"MAINNET_NODE_ADDRESS"`
	MasterAddressKey   string `long:"MASTER_ADDRESS_KEY" env:"MASTER_ADDRESS_KEY"`
	MasterAddress      string `long:"MASTER_ADDRESS" env:"MASTER_ADDRESS"`
}
