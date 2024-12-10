package api

import (
	"github.com/deltadefi-protocol/go-sdk/pkg/utils"
)

type ApiNetwork string

const (
	ApiNetworkPreprod ApiNetwork = "preprod"
	ApiNetworkMainnet ApiNetwork = "mainnet"
)

type ApiConfig struct {
	Network *ApiNetwork `json:"network,omitempty"`
	// TODO: AppWalletKeyType
	SigningKey *string `json:"signingKey,omitempty"`
	utils.AuthHeaders
}
