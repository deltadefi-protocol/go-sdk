package deltadefi

import (
	"fmt"

	rum "github.com/sidan-lab/rum"
	"github.com/sidan-lab/rum/wallet"
)

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

	submitRes, err := d.Order.SubmitPlaceOrderTransactionRequest(&SubmitPlaceOrderTransactionRequest{
		OrderID:  buildRes.OrderID,
		SignedTx: signedTx,
	})
	if err != nil {
		return nil, err
	}
	return submitRes, nil
}

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

	submitRes, err := d.Order.SubmitCancelOrderTransactionRequest(&SubmitCancelOrderTransactionRequest{
		SignedTx: signedTx,
	})
	if err != nil {
		return nil, err
	}
	return submitRes, nil
}
