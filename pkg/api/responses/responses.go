package responses

import (
	"github.com/deltadefi-protocol/go-sdk/pkg/models"
	rmodels "github.com/sidan-lab/rum/models"
)

type SignInResponse struct {
	Token   string `json:"token"`
	IsReady bool   `json:"is_ready"`
}

type BuildSendRefScriptsTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitSendRefScriptsTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}
type SubmitPostOrderTransactionResponse struct {
	Order   *models.OrderJSON `json:"order"`
	TxHexes string            `json:"tx_hexes"`
}

type PostOrderResponse = *SubmitPostOrderTransactionResponse

type DepositRecord struct {
	CreatedAt string          `json:"created_at"`
	Assets    []rmodels.Asset `json:"assets"`
	TxHash    string          `json:"tx_hash"`
}

type GetDepositRecordsResponse []*DepositRecord

type GetOrderRecordResponse struct {
	Orders []*models.OrderJSON `json:"Orders"`
}

type WithdrawalRecord struct {
	CreatedAt string          `json:"created_at"`
	Assets    []rmodels.Asset `json:"assets"`
}

type GetWithdrawalRecordsResponse []*WithdrawalRecord

type AssetBalance struct {
	Asset  string `json:"asset"`
	Free   int64  `json:"free"`
	Locked int64  `json:"locked"`
}

type GetAccountBalanceResponse []*AssetBalance

type GenerateNewAPIKeyResponse struct {
	APIKey string `json:"api_key"`
}

type BuildDepositTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type BuildWithdrawalTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitDepositTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type SubmitWithdrawalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type GetTermsAndConditionResponse struct {
	Value string `json:"value"`
}

type MarketDepth struct {
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

type GetMarketDepthResponse struct {
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

type BuildPlaceOrderTransactionResponse struct {
	OrderID string `json:"order_id"`
	TxHex   string `json:"tx_hex"`
}

type BuildCancelOrderTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitPlaceOrderTransactionResponse struct {
	Order *models.OrderJSON `json:"order"`
}

type SubmitCancelOrderTransactionResponse struct {
	TxHash string `json:"txhash"`
}
