package requests

import (
	"github.com/deltadefi-protocol/go-sdk/pkg/models"
	rmodels "github.com/sidan-lab/rum/models"
)

type SignInRequest struct {
	WalletAddress string `json:"wallet_address"`
	AuthKey       string `json:"auth_key"`
}

// type BuildSendRefScriptsTransactionRequest struct {
// 	InputUTxOs         []rmodels.UTxO  `json:"input_utxos"`
// 	TotalDepositAmount []rmodels.Asset `json:"total_deposit_amount"`
// }

// type SubmitSendRefScriptsTransactionRequest struct {
// 	SignedTx string `json:"signed_tx"`
// }

// type PostOrderRequest = BuildPostOrderTransactionRequest

//	type SubmitDeleteAccountTransactionRequest struct {
//		SignedTx string `json:"signed_tx"`
//	}
type BuildDepositTransactionRequest struct {
	DepositAmount []rmodels.Asset `json:"deposit_amount"`
	InputUtxos    []rmodels.UTxO  `json:"input_utxos"`
}

type BuildWithdrawalTransactionRequest struct {
	WithdrawalAmount []rmodels.Asset `json:"withdrawal_amount"`
}

type SubmitDepositTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

type SubmitWithdrawalTransactionRequest struct {
	SignedTxs []string `json:"signed_txs"`
}

type GetMarketDepthRequest struct {
	Pair models.TradingSymbol `json:"pair"`
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
	Start    *int64   `json:"start,omitempty"` // timestamp
	End      *int64   `json:"end,omitempty"`   // timestamp
}

type BuildPlaceOrderTransactionRequest struct {
	Pair       models.TradingSymbol `json:"pair"`
	Side       models.OrderSide     `json:"side"`
	Type       models.OrderType     `json:"type"`
	Quantity   float64              `json:"quantity"`
	Price      *float64             `json:"price,omitempty"`
	BasisPoint *float64             `json:"basis_point,omitempty"`
}

type SubmitPlaceOrderTransactionRequest struct {
	OrderID  string `json:"order_id"`
	SignedTx string `json:"signed_tx"`
}

type SubmitCancelOrderTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}
