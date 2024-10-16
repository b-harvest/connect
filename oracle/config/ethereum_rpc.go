package config

// ContractConfig holds configuration for a single contract address
type ContractConfig struct {
	Address string   `json:"address"` // Contract address
	Topics  []string `json:"topics"`  // List of topics to filter logs
}
