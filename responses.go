package deltadefi

// GetOperationKeyResponse contains the encrypted operation key and its hash.
type GetOperationKeyResponse struct {
	EncryptedOperationKey string `json:"encrypted_operation_key"`
	OperationKeyHash      string `json:"operation_key_hash"`
}

// BuildSendRefScriptsTransactionResponse contains the transaction hex for reference script sending.
type BuildSendRefScriptsTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

// SubmitSendRefScriptsTransactionResponse contains the transaction hash after submission.
type SubmitSendRefScriptsTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}
type SubmitPostOrderTransactionResponse struct {
	Order   *OrderJSON `json:"order"`
	TxHexes string     `json:"tx_hexes"`
}

type PostOrderResponse = *SubmitPostOrderTransactionResponse

// DepositRecord represents a single deposit transaction record.
type DepositRecord struct {
	CreatedAt string            `json:"created_at"`
	Status    TransactionStatus `json:"status"`
	Assets    []Asset           `json:"assets"`
	TxHash    string            `json:"tx_hash"`
}

// GetDepositRecordsResponse is a collection of deposit transaction records.
type GetDepositRecordsResponse []DepositRecord

// OrderRecordsData represents a single data item containing orders and their execution records.
type OrderRecordsData struct {
	Orders              []OrderJSON              `json:"orders"`
	OrderFillingRecords []OrderFillingRecordJSON `json:"order_filling_records"`
}

// GetOrderRecordsResponse contains paginated order records with metadata.
type GetOrderRecordsResponse struct {
	Data       []OrderRecordsData `json:"data"`
	TotalCount int                `json:"total_count"`
	TotalPage  int                `json:"total_page"`
}

// GetOrderRecordResponse contains details for a single order.
type GetOrderRecordResponse struct {
	OrderJSON OrderJSON `json:"order_json"`
}

// WithdrawalRecord represents a single withdrawal transaction record.
type WithdrawalRecord struct {
	CreatedAt string            `json:"created_at"`
	Status    TransactionStatus `json:"status"`
	Assets    []Asset           `json:"assets"`
}

// GetWithdrawalRecordsResponse is a collection of withdrawal transaction records.
type GetWithdrawalRecordsResponse []WithdrawalRecord

// AssetBalance represents the balance of a specific asset showing free and locked amounts.
type AssetBalance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}

// GetAccountBalanceResponse is a collection of asset balances for the account.
type GetAccountBalanceResponse []AssetBalance

// CreateNewAPIKeyResponse contains a newly generated API key.
type CreateNewAPIKeyResponse struct {
	APIKey string `json:"api_key"`
}

// BuildDepositTransactionResponse contains the transaction hex for deposit operations.
type BuildDepositTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

// BuildWithdrawalTransactionResponse contains the transaction hex for withdrawal operations.
type BuildWithdrawalTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

// BuildTransferalTransactionResponse contains the transaction hex for transfer operations.
type BuildTransferalTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

// SubmitDepositTransactionResponse contains the transaction hash after deposit submission.
type SubmitDepositTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

// SubmitWithdrawalTransactionResponse contains the transaction hash after withdrawal submission.
type SubmitWithdrawalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

// SubmitTransferalTransactionResponse contains the transaction hash after transfer submission.
type SubmitTransferalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

// MarketDepth represents a price level in the order book with price and quantity.
type MarketDepth struct {
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

// GetMarketDepthResponse contains the current order book with bids and asks.
type GetMarketDepthResponse struct {
	Bids []MarketDepth `json:"bids"`
	Asks []MarketDepth `json:"asks"`
}

// GetMarketPriceResponse contains the current market price for a trading pair.
type GetMarketPriceResponse struct {
	Price float64 `json:"price"`
}

// Trade represents a completed trade with price, amount, and metadata.
type Trade struct {
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Side      string  `json:"side"`
	Symbol    string  `json:"symbol"`
	Timestamp string  `json:"timestamp"`
}

// Candlestick represents OHLCV (Open, High, Low, Close, Volume) data for a specific time period.
type Candlestick struct {
	Timestamp int64   `json:"t"`
	Symbol    string  `json:"s"`
	Open      float64 `json:"o"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Close     float64 `json:"c"`
	Volume    float64 `json:"v"`
}

// GetAggregatedPriceResponse is a collection of candlestick data points.
type GetAggregatedPriceResponse []Candlestick

// BuildPlaceOrderTransactionResponse contains the order ID and transaction hex for order placement.
type BuildPlaceOrderTransactionResponse struct {
	OrderID string `json:"order_id"`
	TxHex   string `json:"tx_hex"`
}

// BuildCancelOrderTransactionResponse contains the transaction hex for order cancellation.
type BuildCancelOrderTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

// BuildCancelAllOrdersTransactionResponse contains the transaction hex for order cancellation.
type BuildCancelAllOrdersTransactionResponse struct {
	TxHexes []string `json:"tx_hexes"`
}

// SubmitPlaceOrderTransactionResponse contains the order details after successful submission.
type SubmitPlaceOrderTransactionResponse struct {
	Order OrderJSON `json:"order"`
}

// SubmitCancelOrderTransactionResponse contains the transaction hash after order cancellation.
type SubmitCancelOrderTransactionResponse struct {
	TxHash string `json:"txhash"`
}

// SubmitCancelAllOrdersTransactionResponse contains the transaction hash after order cancellation.
type SubmitCancelAllOrdersTransactionResponse struct {
	CancelledOrderIds []string `json:"cancelled_order_ids"`
}
