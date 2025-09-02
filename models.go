package deltadefi

// OrderStatus represents the various states an order can be in.
type OrderStatus string

const (
	OrderStatusBuilding   OrderStatus = "building"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusOpen       OrderStatus = "open"
	OrderStatusClosed     OrderStatus = "closed"
	OrderStatusFailed     OrderStatus = "failed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

// OrderSide represents whether an order is buying or selling.
type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

// OrderType represents the type of order (market or limit).
type OrderType string

const (
	OrderTypeMarket OrderType = "market"
	OrderTypeLimit  OrderType = "limit"
)

// Symbol represents a trading pair symbol.
type Symbol string

const (
	ADAUSDM Symbol = "ADAUSDM"
)

// OrderJSON represents a complete order with all its details.
type OrderJSON struct {
	OrderID       string                     `json:"order_id"`
	Status        string                     `json:"status"` // Changed from OrderStatus to string
	Symbol        Symbol                     `json:"symbol"`
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

// TransactionStatus represents the various states a transaction can be in.
type TransactionStatus string

const (
	TransactionStatusBuilding         TransactionStatus = "building"
	TransactionStatusHeldForOrder     TransactionStatus = "held_for_order"
	TransactionStatusSubmitted        TransactionStatus = "submitted"
	TransactionStatusSubmissionFailed TransactionStatus = "submission_failed"
	TransactionStatusConfirmed        TransactionStatus = "confirmed"
)

// Asset represents a cryptocurrency asset with its quantity and unit.
type Asset struct {
	Asset     string  `json:"asset"`
	AssetUnit string  `json:"asset_unit"`
	Qty       float64 `json:"qty"`
}

// OrderExecutionRole represents whether the order was a maker or taker in the execution.
type OrderExecutionRole string

const (
	OrderExecutionRoleMaker OrderExecutionRole = "maker"
	OrderExecutionRoleTaker OrderExecutionRole = "taker"
)

// OrderExecutionRecordJSON represents the details of a single order execution/fill.
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

// OrderFillingRecordJSON represents a record of order fill activity.
type OrderFillingRecordJSON struct {
	ExecutionID   string    `json:"execution_id"`
	OrderID       string    `json:"order_id"`
	Status        string    `json:"status"`
	Symbol        Symbol    `json:"symbol"`
	ExecutedQty   string    `json:"executed_qty"`
	Side          OrderSide `json:"side"`
	Type          OrderType `json:"type"`
	FeeCharged    string    `json:"fee_charged"`
	FeeUnit       string    `json:"fee_unit"`
	ExecutedPrice float64   `json:"executed_price"`
	CreatedTime   uint64    `json:"create_time"`
}

// OrderRecordStatus represents the different types of order record queries available.
type OrderRecordStatus string

const (
	OrderRecordStatusOpenOrder      OrderRecordStatus = "openOrder"
	OrderRecordStatusOrderHistory   OrderRecordStatus = "orderHistory"
	OrderRecordStatusTradingHistory OrderRecordStatus = "tradingHistory"
)

// GetOrderRecordRequest contains parameters for querying order records with filtering and pagination.
type GetOrderRecordRequest struct {
	Status OrderRecordStatus `json:"status"`           // Must be either 'openOrder', 'orderHistory', or 'tradingHistory'
	Limit  int               `json:"limit,omitempty"`  // Default is 10, must be between 1 and 250
	Page   int               `json:"page,omitempty"`   // Default is 1, must be between 1 and 1000
	Symbol Symbol            `json:"symbol,omitempty"` // Optional filter by symbol, e.g., ADAUSDM
}
