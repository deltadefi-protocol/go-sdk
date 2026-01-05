package deltadefi

import (
	"fmt"

	"github.com/sidan-lab/rum"
	"github.com/sidan-lab/rum/wallet"
)

// LoadOperationKey loads and decrypts the operation key required for transaction signing.
// This method must be called before performing any transaction operations like placing orders.
//
// Parameters:
//   - passcode: The operation passcode for decrypting the key
//
// Returns:
//   - error: nil on success, error on failure
func (d *DeltaDeFi) LoadOperationKey(passcode string) error {
	res, err := d.Accounts.GetOperationKey()
	if err != nil {
		return err
	}

	// Use proper AES-GCM decryption that matches frontend encryption format
	operationKey, err := rum.DecryptWithCipher(res.EncryptedOperationKey, passcode)
	if err != nil {
		return fmt.Errorf("decryption failed: %w", err)
	}

	operationWallet, err := wallet.NewRootKeyWallet(operationKey, wallet.NewDerivationIndices())
	if err != nil {
		return err
	}

	d.OperationWallet = operationWallet
	return nil
}

// PostOrder is a high-level method for placing an order.
// It handles the complete order flow: building the transaction, signing it, and submitting it.
// The operation wallet must be loaded before calling this method.
//
// Parameters:
//   - data: Order details including symbol, side, type, quantity, and optional price
//
// Returns:
//   - *SubmitPlaceOrderTransactionResponse: Order details and transaction info
//   - error: nil on success, error on failure
func (d *DeltaDeFi) PostOrder(data *BuildPlaceOrderTransactionRequest) (*SubmitPlaceOrderTransactionResponse, error) {
	if d.OperationWallet == nil {
		return nil, fmt.Errorf("operation wallet is not loaded")
	}

	buildRes, err := d.Order.BuildPlaceOrderTransaction(data)
	if err != nil {
		return nil, err
	}

	fmt.Println("Built order, tx hex:", buildRes.TxHex)
	signedTx, err := d.OperationWallet.Signer().SignTransaction(buildRes.TxHex)
	if err != nil {
		return nil, err
	}

	submitRes, err := d.Order.SubmitPlaceOrderTransaction(&SubmitPlaceOrderTransactionRequest{
		OrderID:  buildRes.OrderID,
		SignedTx: signedTx,
	})
	if err != nil {
		return nil, err
	}
	return submitRes, nil
}

// CancelOrder is a high-level method for canceling an existing order.
// This method directly cancels the order without requiring transaction signing.
//
// Parameters:
//   - orderId: The ID of the order to cancel
//
// Returns:
//   - *CancelOrderResponse: Contains the cancelled order ID
//   - error: nil on success, error on failure
func (d *DeltaDeFi) CancelOrder(orderId string) (*CancelOrderResponse, error) {
	return d.Order.CancelOrder(orderId)
}

// CancelAllOrders is a high-level method for canceling all existing orders for a symbol.
// This method directly cancels orders without requiring transaction signing.
//
// Parameters:
//   - symbol: The trading pair symbol to cancel orders for
//
// Returns:
//   - *CancelAllOrdersResponse: Contains the symbol and list of cancelled order IDs
//   - error: nil on success, error on failure
func (d *DeltaDeFi) CancelAllOrders(symbol string) (*CancelAllOrdersResponse, error) {
	return d.Order.CancelAllOrders(symbol)
}
