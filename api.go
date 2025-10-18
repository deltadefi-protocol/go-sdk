package deltadefi

import (
	"fmt"

	rum "github.com/sidan-lab/rum"
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

	operationKey, err := rum.DecryptWithCipher(res.EncryptedOperationKey, passcode)
	if err != nil {
		return err
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
// It handles the complete cancellation flow: building the transaction, signing it, and submitting it.
// The operation wallet must be loaded before calling this method.
//
// Parameters:
//   - orderId: The ID of the order to cancel
//
// Returns:
//   - *SubmitCancelOrderTransactionResponse: Transaction hash of the cancellation
//   - error: nil on success, error on failure
func (d *DeltaDeFi) CancelOrder(orderId string) (*SubmitCancelOrderTransactionResponse, error) {
	if d.OperationWallet == nil {
		return nil, fmt.Errorf("operation wallet is not loaded")
	}

	buildRes, err := d.Order.BuildCancelOrderTransaction(orderId)
	if err != nil {
		return nil, err
	}

	signedTx, err := d.OperationWallet.Signer().SignTransaction(buildRes.TxHex)
	if err != nil {
		return nil, err
	}

	submitRes, err := d.Order.SubmitCancelOrderTransaction(&SubmitCancelOrderTransactionRequest{
		SignedTx: signedTx,
	})
	if err != nil {
		return nil, err
	}
	return submitRes, nil
}
