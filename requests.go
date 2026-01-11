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
	TransferalAmount []rum.Asset    `json:"transferal_amount"`
	ToAddress        string         `json:"to_address"`
	TransferalType   TransferalType `json:"transferal_type,omitempty"`
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
// Note: Either QuoteQuantity or BaseQuantity must be provided, but not both.
type BuildPlaceOrderTransactionRequest struct {
	Symbol                Symbol    `json:"symbol"`
	Side                  OrderSide `json:"side"`
	Type                  OrderType `json:"type"`
	QuoteQuantity         string    `json:"quote_quantity,omitempty"`
	BaseQuantity          string    `json:"base_quantity,omitempty"`
	Price                 string    `json:"price,omitempty"`
	MaxSlippageBasisPoint *string   `json:"max_slippage_basis_point,omitempty"`
	PostOnly              bool      `json:"post_only,omitempty"`
}

// SubmitPlaceOrderTransactionRequest contains the order ID and signed transaction for order submission.
type SubmitPlaceOrderTransactionRequest struct {
	OrderID  string `json:"order_id"`
	SignedTx string `json:"signed_tx"`
}

// CancelAllOrdersRequest contains parameters for canceling all orders for a symbol.
type CancelAllOrdersRequest struct {
	Symbol string `json:"symbol"`
}

// GetOpenOrdersRequest contains parameters for querying open orders.
type GetOpenOrdersRequest struct {
	Symbol string `json:"symbol"`
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
}

// GetTradeOrdersRequest contains parameters for querying trade orders.
type GetTradeOrdersRequest struct {
	Symbol string `json:"symbol"`
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
}

// GetAccountTradesRequest contains parameters for querying account trades.
type GetAccountTradesRequest struct {
	Symbol string `json:"symbol"`
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
}

// GetTransferalRecordsRequest contains parameters for querying transferal records.
type GetTransferalRecordsRequest struct {
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
	Status string `json:"status,omitempty"` // "pending" or "confirmed"
}

// GetDepositRecordsRequest contains parameters for querying deposit records.
type GetDepositRecordsRequest struct {
	Limit int `json:"limit,omitempty"`
	Page  int `json:"page,omitempty"`
}

// GetWithdrawalRecordsRequest contains parameters for querying withdrawal records.
type GetWithdrawalRecordsRequest struct {
	Limit int `json:"limit,omitempty"`
	Page  int `json:"page,omitempty"`
}

// GetAccountBalanceRequest contains parameters for querying account balance.
type GetAccountBalanceRequest struct {
	AssetUnit string `json:"asset_unit,omitempty"`
}

// TransferalType represents the type of transferal transaction.
type TransferalType string

const (
	TransferalTypeDeposit    TransferalType = "deposit"
	TransferalTypeWithdrawal TransferalType = "withdrawal"
)

// BuildRequestTransferalTransactionRequest contains parameters for building a request transferal transaction.
type BuildRequestTransferalTransactionRequest struct {
	TransferalAmount []rum.Asset    `json:"transferal_amount"`
	FromAddress      string         `json:"from_address"`
	TransferalType   TransferalType `json:"transferal_type"`
}

// SubmitRequestTransferalTransactionRequest contains the signed transaction for request transferal submission.
type SubmitRequestTransferalTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

// CreateSpotAccountRequest contains parameters for creating a spot account.
type CreateSpotAccountRequest struct {
	UserId                string `json:"user_id"`
	EncryptedOperationKey string `json:"encrypted_operation_key"`
	OperationKeyHash      string `json:"operation_key_hash"`
	IsScriptOperationKey  *bool  `json:"is_script_operation_key"`
	ReferralCode          string `json:"referral_code,omitempty"`
}

// UpdateSpotAccountRequest contains parameters for updating a spot account.
type UpdateSpotAccountRequest struct {
	UserId                string `json:"user_id"`
	EncryptedOperationKey string `json:"encrypted_operation_key"`
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

// StringPtr returns a pointer to the given string value.
// Useful for setting optional fields in request structures.
func StringPtr(s string) *string {
	return &s
}
