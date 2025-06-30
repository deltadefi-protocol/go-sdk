package deltadefi

func (d *DeltaDeFi) PostOrder(data *BuildPlaceOrderTransactionRequest) (*SubmitPlaceOrderTransactionResponse, error) {
	buildRes, err := d.Order.BuildPlaceOrderTransaction(data)
	if err != nil {
		return nil, err
	}
	// TODO: sign tx
	signedTx := buildRes.TxHex

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
	buildRes, err := d.Order.BuildCancelOrderTransaction(orderId)
	if err != nil {
		return nil, err
	}

	// TODO: sign tx
	signedTx := buildRes.TxHex

	submitRes, err := d.Order.SubmitCancelOrderTransactionRequest(&SubmitCancelOrderTransactionRequest{
		SignedTx: signedTx,
	})
	if err != nil {
		return nil, err
	}
	return submitRes, nil
}
