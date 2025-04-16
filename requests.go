package deltadefi

import (
	"github.com/sidan-lab/rum/models"
)

type SignInRequest struct {
	WalletAddress string `json:"wallet_address"`
	AuthKey       string `json:"auth_key"`
}

// type BuildSendRefScriptsTransactionRequest struct {
// 	InputUTxOs         []models.UTxO  `json:"input_utxos"`
// 	TotalDepositAmount []models.Asset `json:"total_deposit_amount"`
// }

// type SubmitSendRefScriptsTransactionRequest struct {
// 	SignedTx string `json:"signed_tx"`
// }

// type PostOrderRequest = BuildPostOrderTransactionRequest

//	type SubmitDeleteAccountTransactionRequest struct {
//		SignedTx string `json:"signed_tx"`
//	}
type BuildDepositTransactionRequest struct {
	DepositAmount []models.Asset `json:"deposit_amount"`
	InputUtxos    []models.UTxO  `json:"input_utxos"`
}

type BuildWithdrawalTransactionRequest struct {
	WithdrawalAmount []models.Asset `json:"withdrawal_amount"`
}

type SubmitDepositTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

type SubmitWithdrawalTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
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
	Symbol   string   `json:"symbol"`
	Interval Interval `json:"interval"`
	Start    int64    `json:"start"`
	End      int64    `json:"end"`
}

type BuildPlaceOrderTransactionRequest struct {
	Symbol                string    `json:"symbol"`
	Side                  OrderSide `json:"side"`
	Type                  OrderType `json:"type"`
	Quantity              float64   `json:"quantity"`
	Price                 float64   `json:"price"`
	LimitSlippage         bool      `json:"limit_slippage"`
	MaxSlippageBasisPoint bool      `json:"max_slippage_basis_point"`
}

type SubmitPlaceOrderTransactionRequest struct {
	OrderID  string `json:"order_id"`
	SignedTx string `json:"signed_tx"`
}

type SubmitCancelOrderTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}
