package deltadefi

import "github.com/sidan-lab/rum"

// SignInRequest contains credentials for user authentication.
type SignInRequest struct {
	WalletAddress string `json:"wallet_address"`
	AuthKey       string `json:"auth_key"`
}

// type BuildSendRefScriptsTransactionRequest struct {
// 	InputUTxOs         []rum.UTxO  `json:"input_utxos"`
// 	TotalDepositAmount []rum.Asset `json:"total_deposit_amount"`
// }

// type SubmitSendRefScriptsTransactionRequest struct {
// 	SignedTx string `json:"signed_tx"`
// }

// type PostOrderRequest = BuildPostOrderTransactionRequest

//	type SubmitDeleteAccountTransactionRequest struct {
//		SignedTx string `json:"signed_tx"`
//	}
//
// BuildDepositTransactionRequest contains parameters for building a deposit transaction.
type BuildDepositTransactionRequest struct {
	DepositAmount []rum.Asset `json:"deposit_amount"`
	InputUtxos    []rum.UTxO  `json:"input_utxos"`
}

// BuildWithdrawalTransactionRequest contains parameters for building a withdrawal transaction.
type BuildWithdrawalTransactionRequest struct {
	WithdrawalAmount []rum.Asset `json:"withdrawal_amount"`
}

// BuildTransferalTransactionRequest contains parameters for building a transfer transaction.
type BuildTransferalTransactionRequest struct {
	TransferalAmount []rum.Asset `json:"transferal_amount"`
	ToAddress        string      `json:"to_address"`
}

// SubmitDepositTransactionRequest contains the signed transaction for deposit submission.
type SubmitDepositTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

// SubmitWithdrawalTransactionRequest contains the signed transaction for withdrawal submission.
type SubmitWithdrawalTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

// SubmitTransferalTransactionRequest contains the signed transaction for transfer submission.
type SubmitTransferalTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

// Interval represents time intervals for candlestick data.
type Interval string

const (
	Interval5m  Interval = "5m"
	Interval15m Interval = "15m"
	Interval30m Interval = "30m"
	Interval1h  Interval = "1h"
	Interval1d  Interval = "1d"
)

// GetAggregatedPriceRequest contains parameters for retrieving historical price data.
type GetAggregatedPriceRequest struct {
	Symbol   Symbol   `json:"symbol"`
	Interval Interval `json:"interval"`
	Start    int64    `json:"start"`
	End      int64    `json:"end"`
}

// BuildPlaceOrderTransactionRequest contains parameters for building an order placement transaction.
type BuildPlaceOrderTransactionRequest struct {
	Symbol                Symbol    `json:"symbol"`
	Side                  OrderSide `json:"side"`
	Type                  OrderType `json:"type"`
	Quantity              float64   `json:"quantity"`
	Price                 *float64  `json:"price,omitempty"`
	MaxSlippageBasisPoint *int      `json:"max_slippage_basis_point,omitempty"`
	LimitSlippage         *bool     `json:"limit_slippage,omitempty"`
	PostOnly              bool      `json:"post_only,omitempty"`
}

// SubmitPlaceOrderTransactionRequest contains the order ID and signed transaction for order submission.
type SubmitPlaceOrderTransactionRequest struct {
	OrderID  string `json:"order_id"`
	SignedTx string `json:"signed_tx"`
}

// SubmitCancelOrderTransactionRequest contains the signed transaction for order cancellation.
type SubmitCancelOrderTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

// SubmitCancelAllOrdersTransactionRequest contains the signed transaction for order cancellation.
type SubmitCancelAllOrdersTransactionRequest struct {
	SignedTxs string `json:"signed_txs"`
}

// FloatPtr returns a pointer to the given float64 value.
// Useful for setting optional fields in request structures.
func FloatPtr(f float64) *float64 {
	return &f
}

// BoolPtr returns a pointer to the given bool value.
// Useful for setting optional fields in request structures.
func BoolPtr(b bool) *bool {
	return &b
}

// IntPtr returns a pointer to the given int value.
// Useful for setting optional fields in request structures.
func IntPtr(i int) *int {
	return &i
}
