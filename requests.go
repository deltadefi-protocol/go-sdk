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

// type GetMarketDepthRequest struct {
// 	Pair TradingSymbol `json:"pair"`
// }

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
	Start    *int64   `json:"start,omitempty"` // timestamp
	End      *int64   `json:"end,omitempty"`   // timestamp
}

// type BuildPlaceOrderTransactionRequest struct {
// 	Pair       TradingSymbol `json:"pair"`
// 	Side       OrderSide     `json:"side"`
// 	Type       OrderType     `json:"type"`
// 	Quantity   float64       `json:"quantity"`
// 	Price      *float64      `json:"price,omitempty"`
// 	BasisPoint *float64      `json:"basis_point,omitempty"`
// }

type SubmitPlaceOrderTransactionRequest struct {
	OrderID  string `json:"order_id"`
	SignedTx string `json:"signed_tx"`
}

type SubmitCancelOrderTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}
