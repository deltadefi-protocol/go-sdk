package utils

type TradingPair string

const (
	TradingPairADAUSDX TradingPair = "ADAUSDX"
)

type TradingSide string

const (
	TradingSideBuy  TradingSide = "buy"
	TradingSideSell TradingSide = "sell"
)

type TradingType string

const (
	TradingTypeLimit  TradingType = "limit"
	TradingTypeMarket TradingType = "market"
)

type TimeInForce string

const (
	TimeInForceGTC TimeInForce = "GTC"
)
