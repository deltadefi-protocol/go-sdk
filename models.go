package deltadefi

type OrderStatus string

const (
	OrderStatusBuilding   OrderStatus = "building"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusOpen       OrderStatus = "open"
	OrderStatusClosed     OrderStatus = "closed"
	OrderStatusFailed     OrderStatus = "failed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

type OrderType string

const (
	OrderTypeMarket OrderType = "market"
	OrderTypeLimit  OrderType = "limit"
)

type OrderJSON struct {
	OrderID       string                     `json:"order_id"`
	Status        string                     `json:"status"` // Changed from OrderStatus to string
	Symbol        string                     `json:"symbol"`
	OrigQty       string                     `json:"orig_qty"`
	ExecutedQty   string                     `json:"executed_qty"`
	Side          OrderSide                  `json:"side"`
	Price         float64                    `json:"price"`
	Type          OrderType                  `json:"type"`
	FeeCharged    string                     `json:"fee_charged"`
	FeeUnit       string                     `json:"fee_unit"` // Added FeeUnit field
	ExecutedPrice float64                    `json:"executed_price"`
	Slippage      string                     `json:"slippage"`
	CreatedTime   uint64                     `json:"create_time"`     // Changed from CreateTime (int64) to CreatedTime (uint64)
	UpdateTime    uint64                     `json:"update_time"`     // Changed from int64 to uint64
	Fills         []OrderExecutionRecordJSON `json:"fills,omitempty"` // Changed from *[]OrderExecutionRecordJSON to []OrderExecutionRecordJSON
}

// TransactionStatus represents the transaction statuses.
type TransactionStatus string

const (
	TransactionStatusBuilding         TransactionStatus = "building"
	TransactionStatusHeldForOrder     TransactionStatus = "held_for_order"
	TransactionStatusSubmitted        TransactionStatus = "submitted"
	TransactionStatusSubmissionFailed TransactionStatus = "submission_failed"
	TransactionStatusConfirmed        TransactionStatus = "confirmed"
)

type Asset struct {
	Asset     string `json:"asset"`
	AssetUnit string `json:"asset_unit"`
	Qty       int64  `json:"qty"`
}

// OrderExecutionRole represents the role in an order execution.
type OrderExecutionRole string

const (
	OrderExecutionRoleMaker OrderExecutionRole = "maker"
	OrderExecutionRoleTaker OrderExecutionRole = "taker"
)

// OrderExecutionRecordJSON represents an order execution record in JSON format.
type OrderExecutionRecordJSON struct {
	ID                  string             `json:"id"`
	OrderID             string             `json:"order_id"`
	ExecutionPrice      float64            `json:"execution_price"`
	FilledAmount        string             `json:"filled_amount"`
	FeeUnit             string             `json:"fee_unit"`
	FeeAmount           string             `json:"fee_amount"`
	Role                OrderExecutionRole `json:"role"`
	CounterPartyOrderID string             `json:"counter_party_order_id"`
	CreateTime          int64              `json:"create_time"`
}

// OrderFillingRecordJSON represents an order filling record in JSON format.
type OrderFillingRecordJSON struct {
	ExecutionID   string    `json:"execution_id"`
	OrderID       string    `json:"order_id"`
	Status        string    `json:"status"`
	Symbol        string    `json:"symbol"`
	ExecutedQty   string    `json:"executed_qty"`
	Side          OrderSide `json:"side"`
	Type          OrderType `json:"type"`
	FeeCharged    string    `json:"fee_charged"`
	FeeUnit       string    `json:"fee_unit"`
	ExecutedPrice float64   `json:"executed_price"`
	CreatedTime   uint64    `json:"create_time"`
}

// OrderRecordStatus represents the status filter for order records
type OrderRecordStatus string

const (
	OrderRecordStatusOpenOrder      OrderRecordStatus = "openOrder"
	OrderRecordStatusOrderHistory   OrderRecordStatus = "orderHistory"
	OrderRecordStatusTradingHistory OrderRecordStatus = "tradingHistory"
)

// GetOrderRecordRequest represents the request parameters for fetching order records
type GetOrderRecordRequest struct {
	Status OrderRecordStatus `json:"status"`          // Must be either 'openOrder', 'orderHistory', or 'tradingHistory'
	Limit  int               `json:"limit,omitempty"` // Default is 10, must be between 1 and 250
	Page   int               `json:"page,omitempty"`  // Default is 1, must be between 1 and 1000
}
