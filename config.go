package deltadefi

type ApiConfig struct {
	Network           ApiNetwork
	ApiKey            string
	OperationPasscode string
	ProvidedBaseUrl   string
}

type ApiNetwork string

const (
	ApiNetworkDev     ApiNetwork = "dev"
	ApiNetworkStaging ApiNetwork = "staging"
	ApiNetworkMainnet ApiNetwork = "mainnet"
)
