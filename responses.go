package deltadefi

type GetOperationKeyResponse struct {
	EncryptedOperationKey string `json:"encrypted_operation_key"`
	OperationKeyHash      string `json:"operation_key_hash"`
}

type BuildSendRefScriptsTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitSendRefScriptsTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}
type SubmitPostOrderTransactionResponse struct {
	Order   *OrderJSON `json:"order"`
	TxHexes string     `json:"tx_hexes"`
}

type PostOrderResponse = *SubmitPostOrderTransactionResponse

type DepositRecord struct {
	CreatedAt string            `json:"created_at"`
	Status    TransactionStatus `json:"status"`
	Assets    []Asset           `json:"assets"`
	TxHash    string            `json:"tx_hash"`
}

type GetDepositRecordsResponse []DepositRecord

// OrderRecordsData represents a single data item in the order records response
type OrderRecordsData struct {
	Orders              []OrderJSON              `json:"orders"`
	OrderFillingRecords []OrderFillingRecordJSON `json:"order_filling_records"`
}

// GetOrderRecordResponse represents the response from the get order records endpoint.
type GetOrderRecordsResponse struct {
	Data       []OrderRecordsData `json:"data"`
	TotalCount int                `json:"total_count"`
	TotalPage  int                `json:"total_page"`
}

// GetOrderRecordResponse represents the response from the get single order record endpoint.
type GetOrderRecordResponse struct {
	OrderJSON OrderJSON `json:"order_json"`
}

type WithdrawalRecord struct {
	CreatedAt string            `json:"created_at"`
	Status    TransactionStatus `json:"status"`
	Assets    []Asset           `json:"assets"`
}

type GetWithdrawalRecordsResponse []WithdrawalRecord

type AssetBalance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}

type GetAccountBalanceResponse []AssetBalance

type CreateNewAPIKeyResponse struct {
	APIKey string `json:"api_key"`
}

type BuildDepositTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type BuildWithdrawalTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type BuildTransferalTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitDepositTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type SubmitWithdrawalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type SubmitTransferalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
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
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Side      string  `json:"side"`
	Symbol    string  `json:"symbol"`
	Timestamp string  `json:"timestamp"`
}

type Candlestick struct {
	Timestamp int64   `json:"t"`
	Symbol    string  `json:"s"`
	Open      float64 `json:"o"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Close     float64 `json:"c"`
	Volume    float64 `json:"v"`
}

type GetAggregatedPriceResponse []Candlestick

type BuildPlaceOrderTransactionResponse struct {
	OrderID string `json:"order_id"`
	TxHex   string `json:"tx_hex"`
}

type BuildCancelOrderTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitPlaceOrderTransactionResponse struct {
	Order OrderJSON `json:"order"`
}

type SubmitCancelOrderTransactionResponse struct {
	TxHash string `json:"txhash"`
}
