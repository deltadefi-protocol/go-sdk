package utils

import (
	"strconv"

	"github.com/deltadefi-protocol/go-sdk/pkg/api/requests"
	rmodels "github.com/sidan-lab/rum/models"
)

func ConvertUTxO(utxo *rmodels.UTxO) *requests.InputUtxos {
	return &requests.InputUtxos{
		TxHash:  utxo.Input.TxHash,
		TxID:    strconv.Itoa(utxo.Input.OutputIndex),
		Amount:  utxo.Output.Amount,
		Address: utxo.Output.Address,
	}
}

func ConvertUTxOs(utxos []*rmodels.UTxO) []*requests.InputUtxos {
	result := make([]*requests.InputUtxos, len(utxos))
	for i, utxo := range utxos {
		result[i] = ConvertUTxO(utxo)
	}
	return result
}

// TODO: add TxInParameter in rum
// func ConvertTxInParameter(txIn *TxInParameter) *InputUtxos {
// 	return &InputUtxos{
// 		TxHash:  txIn.TxHash,
// 		TxID:    strconv.Itoa(txIn.TxIndex),
// 		Amount:  txIn.Amount,
// 		Address: txIn.Address,
// 	}
// }
