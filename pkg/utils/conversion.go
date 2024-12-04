package utils

import (
	"strconv"

	rmodels "github.com/sidan-lab/rum/models"
)

func ConvertUTxO(utxo *rmodels.UTxO) *InputUtxos {
	return &InputUtxos{
		TxHash:  utxo.Input.TxHash,
		TxID:    strconv.Itoa(utxo.Input.OutputIndex),
		Amount:  utxo.Output.Amount,
		Address: utxo.Output.Address,
	}
}

func ConvertUTxOs(utxos []*rmodels.UTxO) []*InputUtxos {
	result := make([]*InputUtxos, len(utxos))
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
