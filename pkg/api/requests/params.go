package requests

import (
	"github.com/sidan-lab/rum/models"
)

type InputUtxos struct {
	TxHash  string       `json:"tx_hash"`
	TxID    string       `json:"tx_id"`
	Amount  models.Asset `json:"amount"`
	Address string       `json:"address"`
}
