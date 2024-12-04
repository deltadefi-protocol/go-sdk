package responses

import (
	"github.com/deltadefi-protocol/go-sdk/pkg/models"
	rmodels "github.com/sidan-lab/rum/models"
)

type SignInResponse struct {
	Token   string `json:"token"`
	IsReady bool   `json:"is_ready"`
}

type Balance struct {
	Total                  map[string]int64 `json:"total"`
	AvailableForTrade      map[string]int64 `json:"available_for_trade"`
	AvailableForWithdrawal map[string]int64 `json:"available_for_withdrawal"`
	HeldForOrder           map[string]int64 `json:"held_for_order"`
	SpendingSettling       map[string]int64 `json:"spending_settling"`
	DepositingSettling     map[string]int64 `json:"depositing_settling"`
}

type GetBalanceResponse struct {
	Balance *Balance `json:"balance"`
}

type GetOrdersResponse struct {
	Orders []*models.Order `json:"orders"`
}

type BuildSendRefScriptsTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitSendRefScriptsTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type BuildDepositTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitDepositTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type BuildWithdrawalTransactionResponse struct {
	TxHexes []string `json:"tx_hexes"`
}

type SubmitWithdrawalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type BuildPostOrderTransactionResponse struct {
	OrderID    string   `json:"order_id"`
	ChainedTxs []string `json:"chained_txs"`
	TxHexes    []string `json:"tx_hexes"`
}

type SubmitPostOrderTransactionResponse struct {
	Order   *models.Order `json:"order"`
	TxHexes []string      `json:"tx_hexes"`
}

type PostOrderResponse = SubmitPostOrderTransactionResponse

type CancelOrderResponse struct {
	Message string `json:"message"`
}

type MarketDepth struct {
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

type GetDepthResponse struct {
	Bids []MarketDepth `json:"bids"`
	Asks []MarketDepth `json:"asks"`
}

type GetMarketPriceResponse struct {
	Price float64 `json:"price"`
}

type Trade struct {
	Time   string  `json:"time"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

type GetAggregatedPriceResponse []*Trade

type GetAccountInfoResponse struct {
	APIKey        string `json:"api_key"`
	APILimit      string `json:"api_limit"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	WalletAddress string `json:"wallet_address"`
	IsReady       bool   `json:"is_ready"`
}

type GetNewApiKeyResponse struct {
	APIKey string `json:"api_key"`
}

type BuildDeleteAccountTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitDeleteAccountTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type GetDepositInfoResponse struct {
	TotalDeposit     *DepositInfo `json:"total_deposit"`
	SuggestedDeposit *DepositInfo `json:"suggested_deposit"`
}

type DepositInfo struct {
	Amount             []*rmodels.Asset `json:"amount"`
	PostDepositBalance *Balance         `json:"post_deposit_balance"`
}
