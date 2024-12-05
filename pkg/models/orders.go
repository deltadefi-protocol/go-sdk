package models

type TradingSymbol string

const (
	ADAUSDX TradingSymbol = "ADAUSDX"
)

type OrderStatus string

const (
	OrderStatusBuilding OrderStatus = "building"
	OrderStatusOpen     OrderStatus = "open"
	OrderStatusClosed   OrderStatus = "closed"
	OrderStatusFailed   OrderStatus = "failed"
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
	OrderID       string        `json:"order_id"`
	Status        OrderStatus   `json:"status"`
	Symbol        TradingSymbol `json:"symbol"`
	OrigQty       string        `json:"orig_qty"`
	ExecutedQty   string        `json:"executed_qty"`
	Side          OrderSide     `json:"side"`
	Price         string        `json:"price"`
	Type          OrderType     `json:"type"`
	FeeAmount     float64       `json:"fee_amount"`
	ExecutedPrice float64       `json:"executed_price"`
	Slippage      string        `json:"slippage"`
	CreateTime    int64         `json:"create_time"`
	UpdateTime    int64         `json:"update_time"`
}
