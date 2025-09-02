package deltadefi

// ApiConfig contains configuration settings for the DeltaDeFi client.
type ApiConfig struct {
	// Network specifies which environment to connect to (dev, staging, mainnet)
	Network           ApiNetwork
	// ApiKey is the authentication key for API access
	ApiKey            string
	// OperationPasscode is used for decrypting operation keys for transaction signing
	OperationPasscode string
	// ProvidedBaseUrl allows overriding the default API base URL (optional)
	ProvidedBaseUrl   string
}

// ApiNetwork represents the different network environments available.
type ApiNetwork string

const (
	// ApiNetworkDev connects to the development environment
	ApiNetworkDev     ApiNetwork = "dev"
	// ApiNetworkStaging connects to the staging environment for testing
	ApiNetworkStaging ApiNetwork = "staging"
	// ApiNetworkMainnet connects to the production mainnet environment
	ApiNetworkMainnet ApiNetwork = "mainnet"
)
