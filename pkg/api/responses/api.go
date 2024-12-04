package responses

type ApiNetwork string

const (
	ApiNetworkPreprod ApiNetwork = "preprod"
	ApiNetworkMainnet ApiNetwork = "mainnet"
)

type ApiConfig struct {
	Network *ApiNetwork `json:"network,omitempty"`
	//TODO AppWalletKeyType
	SigningKey int `json:"signingKey,omitempty"`
	AuthHeaders
}
