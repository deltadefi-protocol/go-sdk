# DeltaDeFi Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/deltadefi-protocol/go-sdk.svg)](https://pkg.go.dev/github.com/deltadefi-protocol/go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/deltadefi-protocol/go-sdk)](https://goreportcard.com/report/github.com/deltadefi-protocol/go-sdk)

Official Go SDK for DeltaDeFi, a decentralized finance protocol built on Cardano. This SDK provides a simple and intuitive interface to interact with the DeltaDeFi API for trading, account management, and market data operations.

## Features

- ✅ **Trading Operations** - Place, cancel, and manage orders
- ✅ **Account Management** - Deposits, withdrawals, balance queries
- ✅ **Market Data** - Real-time prices and historical data
- ✅ **Wallet Integration** - Secure transaction signing with Cardano wallets
- ✅ **Multiple Networks** - Support for mainnet, staging, and development environments
- ✅ **Type Safety** - Full Go type definitions for all API endpoints

## Installation

```bash
go get github.com/deltadefi-protocol/go-sdk
```

## Requirements

- Go 1.21 or higher
- Valid DeltaDeFi API key
- Operation passcode for transaction signing

## Quick Start

### Basic Setup

```go
package main

import (
    "fmt"
    "log"

    deltadefi "github.com/deltadefi-protocol/go-sdk"
)

func main() {
    // Initialize the client
    config := deltadefi.ApiConfig{
        Network:           deltadefi.ApiNetworkStaging, // or ApiNetworkMainnet
        ApiKey:            "your-api-key-here",
        OperationPasscode: "your-operation-passcode",
    }

    client := deltadefi.NewDeltaDeFi(config)

    // Load operation key for transaction signing
    err := client.LoadOperationKey(config.OperationPasscode)
    if err != nil {
        log.Fatal("Failed to load operation key:", err)
    }

    // Get account balance
    balance, err := client.Accounts.GetAccountBalance()
    if err != nil {
        log.Fatal("Failed to get balance:", err)
    }

    fmt.Printf("Account Balance: %+v\n", balance)
}
```

### Environment Configuration

Create a `.env` file based on `.env.example`:

```env
# Network configuration
NETWORK=staging  # or "mainnet" for production

# API credentials (optional for testing)
MAESTRO_API_KEY=your-maestro-api-key
BLOCKFROST_PROJECT_ID=your-blockfrost-project-id
```

## API Reference

### Client Initialization

#### `NewDeltaDeFi(config ApiConfig) *DeltaDeFi`

Creates a new DeltaDeFi client instance.

**Parameters:**

- `config`: Configuration object containing network, API key, and operation passcode

**Supported Networks:**

- `ApiNetworkDev`: Development environment
- `ApiNetworkStaging`: Staging environment
- `ApiNetworkMainnet`: Production environment

### Authentication

#### `LoadOperationKey(passcode string) error`

Loads the encrypted operation key required for transaction signing.

**Parameters:**

- `passcode`: Your operation passcode

**Returns:** Error if key loading fails

## Account Management

### Get Account Balance

```go
balance, err := client.Accounts.GetAccountBalance()
```

**Response:** `[]AssetBalance` - Array of asset balances with free and locked amounts

### Get Operation Key

```go
operationKey, err := client.Accounts.GetOperationKey()
```

**Response:** `GetOperationKeyResponse` - Encrypted operation key and hash

### Create New API Key

```go
apiKey, err := client.Accounts.CreateNewAPIKey()
```

**Response:** `CreateNewAPIKeyResponse` - New API key string

### Transaction Records

#### Get Deposit Records

```go
deposits, err := client.Accounts.GetDepositRecords()
```

**Response:** `[]DepositRecord` - Array of deposit transaction records

#### Get Withdrawal Records

```go
withdrawals, err := client.Accounts.GetWithdrawalRecords()
```

**Response:** `[]WithdrawalRecord` - Array of withdrawal transaction records

#### Get Order Records

```go
request := &deltadefi.GetOrderRecordRequest{
    Status: deltadefi.OrderRecordStatusOpenOrder, // or OrderRecordStatusOrderHistory, OrderRecordStatusTradingHistory
    Limit:  10,   // Optional: 1-250, default 10
    Page:   1,    // Optional: 1-1000, default 1
    Symbol: deltadefi.ADAUSDM, // Optional: filter by trading pair
}

orders, err := client.Accounts.GetOrderRecords(request)
```

**Response:** `GetOrderRecordsResponse` - Paginated order records with total count

#### Get Single Order Record

```go
order, err := client.Accounts.GetOrderRecord("order-id-here")
```

**Response:** `GetOrderRecordResponse` - Single order details

### Transaction Building and Submission

#### Deposit Transaction

```go
// Build deposit transaction
buildRequest := &deltadefi.BuildDepositTransactionRequest{
    DepositAmount: []rum.Asset{{Asset: "lovelace", Qty: 100.0}},
    InputUtxos:    utxos, // Your input UTXOs
}

buildResponse, err := client.Accounts.BuildDepositTransaction(buildRequest)
if err != nil {
    log.Fatal(err)
}

// Sign and submit
signedTx, err := client.OperationWallet.Signer().SignTransaction(buildResponse.TxHex)
if err != nil {
    log.Fatal(err)
}

submitRequest := &deltadefi.SubmitDepositTransactionRequest{
    SignedTx: signedTx,
}

result, err := client.Accounts.SubmitDepositTransaction(submitRequest)
```

#### Withdrawal Transaction

```go
// Build withdrawal transaction
buildRequest := &deltadefi.BuildWithdrawalTransactionRequest{
    WithdrawalAmount: []rum.Asset{{Asset: "lovelace", Qty: 50.0}},
}

buildResponse, err := client.Accounts.BuildWithdrawalTransaction(buildRequest)
// Sign and submit similar to deposit
```

#### Transfer Transaction

```go
// Build transfer transaction
buildRequest := &deltadefi.BuildTransferalTransactionRequest{
    TransferalAmount: []rum.Asset{{Asset: "lovelace", Qty: 25.0}},
    ToAddress:        "addr1...", // Destination address
}

buildResponse, err := client.Accounts.BuildTransferalTransaction(buildRequest)
// Sign and submit similar to deposit
```

## Market Data

### Get Market Price

```go
price, err := client.Market.GetMarketPrice("ADAUSDM")
```

**Response:** `GetMarketPriceResponse` - Current market price

### Get Aggregated Price Data (Candlesticks)

```go
request := &deltadefi.GetAggregatedPriceRequest{
    Symbol:   deltadefi.ADAUSDM,
    Interval: deltadefi.Interval1h, // 5m, 15m, 30m, 1h, 1d
    Start:    1640995200, // Unix timestamp
    End:      1641081600, // Unix timestamp
}

candlesticks, err := client.Market.GetAggregatedPrice(request)
```

**Response:** `[]Candlestick` - Array of OHLCV data points

## Order Management

### Place Order (High-level)

```go
// Market order
orderRequest := &deltadefi.BuildPlaceOrderTransactionRequest{
    Symbol:   deltadefi.ADAUSDM,
    Side:     deltadefi.OrderSideBuy,
    Type:     deltadefi.OrderTypeMarket,
    Quantity: 100.0,
}

result, err := client.PostOrder(orderRequest)
```

### Place Order (Low-level)

```go
// Build order transaction
buildRequest := &deltadefi.BuildPlaceOrderTransactionRequest{
    Symbol:                deltadefi.ADAUSDM,
    Side:                  deltadefi.OrderSideBuy,
    Type:                  deltadefi.OrderTypeLimit,
    Quantity:              100.0,
    Price:                 deltadefi.FloatPtr(1.25), // Required for limit orders
    LimitSlippage:         deltadefi.BoolPtr(true),  // Optional
    MaxSlippageBasisPoint: deltadefi.IntPtr(50),     // Optional: 0.5%
}

buildResponse, err := client.Order.BuildPlaceOrderTransaction(buildRequest)
if err != nil {
    log.Fatal(err)
}

// Sign transaction
signedTx, err := client.OperationWallet.Signer().SignTransaction(buildResponse.TxHex)
if err != nil {
    log.Fatal(err)
}

// Submit order
submitRequest := &deltadefi.SubmitPlaceOrderTransactionRequest{
    OrderID:  buildResponse.OrderID,
    SignedTx: signedTx,
}

result, err := client.Order.SubmitPlaceOrderTransactionRequest(submitRequest)
```

### Cancel Order

```go
// High-level cancel
result, err := client.CancelOrder("order-id-here")

// Or low-level cancel
buildResponse, err := client.Order.BuildCancelOrderTransaction("order-id-here")
// Sign and submit similar to place order
```

## Data Types

### Order Types and Status

```go
// Order sides
deltadefi.OrderSideBuy
deltadefi.OrderSideSell

// Order types
deltadefi.OrderTypeMarket
deltadefi.OrderTypeLimit

// Order status
deltadefi.OrderStatusOpen
deltadefi.OrderStatusClosed
deltadefi.OrderStatusCancelled
deltadefi.OrderStatusFailed

// Order record status filters
deltadefi.OrderRecordStatusOpenOrder      // Active orders
deltadefi.OrderRecordStatusOrderHistory   // Historical orders
deltadefi.OrderRecordStatusTradingHistory // Trade executions
```

### Supported Trading Pairs

```go
deltadefi.ADAUSDM // ADA/USDM pair
```

### Time Intervals

```go
deltadefi.Interval5m   // 5 minutes
deltadefi.Interval15m  // 15 minutes
deltadefi.Interval30m  // 30 minutes
deltadefi.Interval1h   // 1 hour
deltadefi.Interval1d   // 1 day
```

## Error Handling

The SDK returns standard Go errors. Always check for errors in production code:

```go
balance, err := client.Accounts.GetAccountBalance()
if err != nil {
    // Handle error appropriately
    log.Printf("API error: %v", err)
    return
}

// Process successful response
for _, asset := range *balance {
    fmt.Printf("Asset: %s, Free: %.2f, Locked: %.2f\n",
        asset.Asset, asset.Free, asset.Locked)
}
```

## Helper Functions

The SDK provides helper functions for creating optional pointer values:

```go
// For optional float64 fields
price := deltadefi.FloatPtr(1.25)

// For optional bool fields
limitSlippage := deltadefi.BoolPtr(true)

// For optional int fields
maxSlippage := deltadefi.IntPtr(50)
```

## Security Best Practices

1. **API Key Management**: Never commit API keys to version control
2. **Environment Variables**: Use environment variables or secure vaults for credentials
3. **Network Selection**: Use appropriate network (staging for testing, mainnet for production)
4. **Error Handling**: Always handle errors and validate responses
5. **Passcode Security**: Store operation passcodes securely and never log them

## Examples

See the [sdks-demo/go](https://github.com/deltadefi-protocol/sdks-demo/tree/main/go) for complete working examples

## Dependencies

- [github.com/sidan-lab/rum](https://github.com/sidan-lab/rum) - Cardano wallet and transaction utilities

## License

This project is licensed under the Apache2.0 License - see the LICENSE file for details.

## Support

- 📧 Email: <contact@deltadefi.io>
- 📚 Documentation: [docs.deltadefi.io](https://docs.deltadefi.io)
- 🐛 Issues: [GitHub Issues](https://github.com/deltadefi-protocol/go-sdk/issues)
- Telegram community: [Telegram Community](https://t.me/deltadefi_community)
- Discord community: [Dicord Community](https://discord.gg/55Z25r3QfC)
