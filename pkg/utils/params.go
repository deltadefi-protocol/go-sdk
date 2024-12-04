package utils

import (
	rmodels "github.com/sidan-lab/rum/models"
)

type InputUtxos struct {
	TxHash  string          `json:"tx_hash"`
	TxID    string          `json:"tx_id"`
	Amount  []rmodels.Asset `json:"amount"`
	Address string          `json:"address"`
}
