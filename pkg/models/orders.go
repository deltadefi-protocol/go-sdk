package models

import (
	"github.com/deltadefi-protocol/go-sdk/pkg/utils"
)

type OrderStatus string

const (
	OrderStatusBuild         OrderStatus = "build"
	OrderStatusPending       OrderStatus = "pending"
	OrderStatusOpen          OrderStatus = "open"
	OrderStatusCancelled     OrderStatus = "cancelled"
	OrderStatusPartialFilled OrderStatus = "partial_filled"
	OrderStatusPendingSettle OrderStatus = "pending_settle"
	OrderStatusFullyFilled   OrderStatus = "fully_filled"
	OrderStatusFailed        OrderStatus = "failed"
)

type Order struct {
	Pair        utils.TradingPair `json:"pair"`
	OrderID     string            `json:"order_id"`
	Price       string            `json:"price"`
	Slippage    float64           `json:"slippage"`
	OrigQty     string            `json:"orig_qty"`
	ExecutedQty string            `json:"executed_qty"`
	SettlingQty string            `json:"settling_qty"`
	Status      OrderStatus       `json:"status"`
	TimeInForce utils.TimeInForce `json:"time_in_force"`
	ExpiryTime  int64             `json:"expiry_time"`
	Type        utils.TradingType `json:"type"`
	Side        utils.TradingSide `json:"side"`
	CreateTime  int64             `json:"create_time"`
	UpdateTime  int64             `json:"update_time"`
}
