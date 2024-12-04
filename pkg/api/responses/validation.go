package responses

import (
	"github.com/sidan-lab/rum/models"
)

/**
 * DeltaDeFiOrderInfo is a type that represents the information of a DeltaDeFi order.
 * @property {Asset[]} assetsToPay - The assets that are to be paid from orders in current transaction.
 * @property {Asset[]} assetsToReturn - The assets that are to be received from orders in current transaction.
 * @property {string} txFee - The transaction fee.
 * @property {string} tradingFee - The trading fee.
 */
type DeltaDeFiOrderInfo struct {
	AssetsToPay    []models.Asset `json:"assetsToPay"`
	AssetsToReturn []models.Asset `json:"assetsToReturn"`
	TxFee          string         `json:"txFee"`
	TradingFee     string         `json:"tradingFee"`
}

/**
 * DeltaDeFiTxInfo is a type that represents the information of a DeltaDeFi transaction.
 * @property {Asset[]} accountInput - The assets that are input from the account.
 * @property {Asset[]} accountOutput - The assets that are output to the account.
 * @property {Asset[]} dexInput - The assets that are input from the DEX.
 * @property {Asset[]} dexOutput - The assets that are output to the DEX.
 * @property {string} txFee - The transaction fee.
 * @property {string} tradingFee - The trading fee.
 */
type DeltaDeFiTxInfo struct {
	AccountInput  []models.Asset       `json:"accountInput"`
	AccountOutput []models.Asset       `json:"accountOutput"`
	DexInput      []DeltaDeFiOrderInfo `json:"dexInput"`
	DexOutput     []DeltaDeFiOrderInfo `json:"dexOutput"`
	TxFee         string               `json:"txFee"`
	TradingFee    string               `json:"tradingFee"`
}
