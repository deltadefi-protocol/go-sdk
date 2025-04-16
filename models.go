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
	OrderID       string                      `json:"order_id"`
	Status        OrderStatus                 `json:"status"`
	Symbol        string                      `json:"symbol"`
	OrigQty       string                      `json:"orig_qty"`
	ExecutedQty   string                      `json:"executed_qty"`
	Side          OrderSide                   `json:"side"`
	Price         float64                     `json:"price"`
	Type          OrderType                   `json:"type"`
	FeeCharged    string                      `json:"fee_charged"`
	ExecutedPrice float64                     `json:"executed_price"`
	Slippage      string                      `json:"slippage"`
	CreateTime    int64                       `json:"create_time"`
	UpdateTime    int64                       `json:"update_time"`
	Fills         *[]OrderExecutionRecordJSON `json:"fills"`
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
