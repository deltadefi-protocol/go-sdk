package requests

import (
	"github.com/deltadefi-protocol/go-sdk/pkg/utils"
	"github.com/sidan-lab/rum/models"
)

type SignInRequest struct {
	WalletAddress string `json:"wallet_address"`
	AuthKey       string `json:"auth_key"`
}

type BuildSendRefScriptsTransactionRequest struct {
	InputUTxOs         []models.UTxO  `json:"input_utxos"`
	TotalDepositAmount []models.Asset `json:"total_deposit_amount"`
}

type SubmitSendRefScriptsTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

type BuildDepositTransactionRequest struct {
	DepositAmount []models.Asset `json:"deposit_amount"`
	InputUTxOs    []models.UTxO  `json:"input_utxos"`
}

type SubmitDepositTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

type BuildWithdrawalTransactionRequest struct {
	WithdrawalAmount []models.Asset `json:"withdrawal_amount"`
	InputUTxOs       []models.UTxO  `json:"input_utxos"`
}

type SubmitWithdrawalTransactionRequest struct {
	SignedTxs []string `json:"signed_txs"`
}

type BuildPostOrderTransactionRequest struct {
	Pair       utils.TradingPair `json:"pair"`
	Side       utils.TradingSide `json:"side"`
	Type       utils.TradingType `json:"type"`
	Quantity   float64           `json:"quantity"`
	Price      *float64          `json:"price,omitempty"`
	BasisPoint *float64          `json:"basis_point,omitempty"`
}

type SubmitPostOrderTransactionRequest struct {
	OrderID   string   `json:"order_id"`
	SignedTxs []string `json:"signed_txs"`
}

type PostOrderRequest = BuildPostOrderTransactionRequest

type GetDepthRequest struct {
	Pair string `json:"pair"`
}

type GetMarketPriceRequest struct {
	Pair string `json:"pair"`
}

type Interval string

const (
	Interval15m Interval = "15m"
	Interval30m Interval = "30m"
	Interval1h  Interval = "1h"
	Interval1d  Interval = "1d"
	Interval1w  Interval = "1w"
	Interval1M  Interval = "1M"
)

type GetAggregatedPriceRequest struct {
	Pair     string   `json:"pair"`
	Interval Interval `json:"interval"`
	Start    *int64   `json:"start,omitempty"`
	End      *int64   `json:"end,omitempty"`
}

type SubmitDeleteAccountTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

type GetDepositInfoRequest struct {
	TotalDepositAmount []models.Asset `json:"total_deposit_amount"`
}
