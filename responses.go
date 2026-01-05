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

// SubmitPlaceOrderTransactionResponse contains the order details after successful submission.
type SubmitPlaceOrderTransactionResponse struct {
	Order OrderJSON `json:"order"`
}

// CancelOrderResponse contains the order ID of the cancelled order.
type CancelOrderResponse struct {
	OrderId string `json:"order_id"`
}

// CancelAllOrdersResponse contains details of all cancelled orders.
type CancelAllOrdersResponse struct {
	Symbol   string   `json:"symbol"`
	OrderIds []string `json:"order_ids"`
}

// GetMaxDepositResponse contains the maximum deposit amount.
type GetMaxDepositResponse struct {
	MaxDeposit string `json:"max_deposit"`
}

// GetAPIKeyResponse contains the API key and creation timestamp.
type GetAPIKeyResponse struct {
	ApiKey    string `json:"api_key"`
	CreatedAt string `json:"created_at"`
}

// GetSpotAccountResponse contains spot account details.
type GetSpotAccountResponse struct {
	AccountID             string `json:"account_id"`
	AccountType           string `json:"account_type"`
	EncryptedOperationKey string `json:"encrypted_operation_key"`
	OperationKeyHash      string `json:"operation_key_hash"`
	CreatedAt             string `json:"created_at"`
}

// CreateSpotAccountResponse contains the created spot account details.
type CreateSpotAccountResponse struct {
	AccountID             string `json:"account_id"`
	AccountType           string `json:"account_type"`
	EncryptedOperationKey string `json:"encrypted_operation_key"`
	OperationKeyHash      string `json:"operation_key_hash"`
	CreatedAt             string `json:"created_at"`
}

// UpdateSpotAccountResponse contains the updated spot account details.
type UpdateSpotAccountResponse struct {
	AccountID             string `json:"account_id"`
	AccountType           string `json:"account_type"`
	EncryptedOperationKey string `json:"encrypted_operation_key"`
	OperationKeyHash      string `json:"operation_key_hash"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
}

// TransferalRecord represents a single transferal transaction record.
type TransferalRecord struct {
	CreatedAt      string            `json:"created_at"`
	Status         TransactionStatus `json:"status"`
	Assets         []Asset           `json:"assets"`
	TxHash         string            `json:"tx_hash"`
	ToAddress      string            `json:"to_address,omitempty"`
	FromAddress    string            `json:"from_address,omitempty"`
	TransferalType string            `json:"transferal_type,omitempty"`
}

// GetTransferalRecordsResponse is a collection of transferal transaction records.
type GetTransferalRecordsResponse []TransferalRecord

// GetTransferalRecordByTxHashResponse contains a single transferal record.
type GetTransferalRecordByTxHashResponse struct {
	TransferalRecord TransferalRecord `json:"transferal_record"`
}

// BuildRequestTransferalTransactionResponse contains the transaction hex for request transferal.
type BuildRequestTransferalTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

// SubmitRequestTransferalTransactionResponse contains the transaction hash after request transferal submission.
type SubmitRequestTransferalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

// OrderResponse represents an order with quantities in human-readable format (from Espresso develop).
type OrderResponse struct {
	ID                    string                         `json:"id"`
	AccountID             string                         `json:"account_id"`
	ActiveOrderUtxoID     *string                        `json:"active_order_utxo_id,omitempty"`
	Status                string                         `json:"status"`
	Symbol                Symbol                         `json:"symbol"`
	BaseQty               string                         `json:"base_qty"`
	QuoteQty              string                         `json:"quote_qty"`
	Side                  OrderSide                      `json:"side"`
	Price                 string                         `json:"price"`
	Type                  OrderType                      `json:"type"`
	SlippageBp            *uint64                        `json:"slippage_bp,omitempty"`
	MarketOrderLimitPrice *string                        `json:"market_order_limit_price,omitempty"`
	LockedBaseQty         string                         `json:"locked_base_qty"`
	LockedQuoteQty        string                         `json:"locked_quote_qty"`
	ExecutedBaseQty       string                         `json:"executed_base_qty"`
	ExecutedQuoteQty      string                         `json:"executed_quote_qty"`
	ObOpenOrderBaseQty    string                         `json:"ob_open_order_base_qty"`
	CommissionUnit        string                         `json:"commission_unit"`
	Commission            string                         `json:"commission"`
	CommissionRateBp      uint64                         `json:"commission_rate_bp"`
	ExecutedPrice         string                         `json:"executed_price"`
	CreatedAt             string                         `json:"created_at"`
	UpdatedAt             string                         `json:"updated_at"`
	OrderExecutionRecords []OrderExecutionRecordResponse `json:"order_execution_records,omitempty"`
}

// OrderExecutionRecordResponse represents a trade execution with quantities in human-readable format.
type OrderExecutionRecordResponse struct {
	ID                  string `json:"id"`
	OrderID             string `json:"order_id"`
	AccountID           string `json:"account_id"`
	ExecutionPrice      string `json:"execution_price"`
	FilledBaseQty       string `json:"filled_base_qty"`
	FilledQuoteQty      string `json:"filled_quote_qty"`
	CommissionUnit      string `json:"commission_unit"`
	Commission          string `json:"commission"`
	Role                string `json:"role"`
	CounterPartyOrderID string `json:"counter_party_order_id"`
	CreatedAt           string `json:"created_at"`
}

// GetOpenOrdersResponse contains paginated open orders.
type GetOpenOrdersResponse struct {
	Data       []OrderResponse `json:"data"`
	TotalCount int             `json:"total_count"`
	TotalPage  int             `json:"total_page"`
}

// GetTradeOrdersResponse contains paginated trade orders.
type GetTradeOrdersResponse struct {
	Data       []OrderResponse `json:"data"`
	TotalCount int             `json:"total_count"`
	TotalPage  int             `json:"total_page"`
}

// GetAccountTradesResponse contains paginated account trades.
type GetAccountTradesResponse struct {
	Data       []OrderExecutionRecordResponse `json:"data"`
	TotalCount int                            `json:"total_count"`
	TotalPage  int                            `json:"total_page"`
}
