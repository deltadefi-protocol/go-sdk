package utils

import (
	"strconv"

	rmodels "github.com/sidan-lab/rum/models"
)

type Value struct {
	Value map[string]int64 `json:"value"`
}

func NewValue() *Value {
	return &Value{
		Value: make(map[string]int64),
	}
}

/**
 * Add an asset to the Value class's value record. If an asset with the same unit already exists in the value record, the quantity of the
 * existing asset will be increased by the quantity of the new asset. If no such asset exists, the new asset will be added to the value record.
 * Implementation:
 * 1. Check if the unit of the asset already exists in the value record.
 * 2. If the unit exists, add the new quantity to the existing quantity.
 * 3. If the unit does not exist, add the unti to the object.
 * 4. Return the Value class instance.
 * @param asset
 * @returns this
 */
func (v *Value) AddAsset(asset *rmodels.Asset) *Value {
	quantity, err := strconv.ParseInt(asset.Quantity, 10, 64)
	if err != nil {
		// TODO: Handle error
	}

	if v.Value[asset.Unit] != 0 {
		v.Value[asset.Unit] += quantity
	} else {
		v.Value[asset.Unit] = quantity
	}
	return v
}

/**
 * Add an array of assets to the Value class's value record. If an asset with the same unit already exists in the value record, the quantity of the
 * existing asset will be increased by the quantity of the new asset. If no such asset exists, the new assets under the array of assets will be added to the value record.
 * Implementation:
 * 1. Iterate over each asset in the 'assets' array.
 * 2. For each asset, check if the unit of the asset already exists in the value record.
 * 3. If the unit exists, add the new quantity to the existing quantity.
 * 4. If the unit does not exist, add the unti to the object.
 * 5. Return the Value class instance.
 * @param assets
 * @returns this
 */
func (v *Value) AddAssets(assets []*rmodels.Asset) *Value {
	for _, asset := range assets {
		v.AddAsset(asset)
	}
	return v
}

/**
 * Substract an asset from the Value class's value record. If an asset with the same unit already exists in the value record, the quantity of the
 * existing asset will be decreased by the quantity of the new asset. If no such asset exists, an error message should be printed.
 * Implementation:
 * 1. Check if the unit of the asset already exists in the value record.
 * 2. If the unit exists, subtract the new quantity from the existing quantity.
 * 3. If the unit does not exist, print an error message.
 * @param asset
 * @returns this
 */
func (v *Value) NegateAsset(asset *rmodels.Asset) *Value {
	unit := asset.Unit
	quantity, err := strconv.ParseInt(asset.Quantity, 10, 64)
	if err != nil {
		// TODO: Handle error
	}

	currentQuantity, exists := v.Value[unit]
	if !exists {
		// TODO: Handle error
	}
	newQuantity := currentQuantity - quantity
	if newQuantity == 0 {
		delete(v.Value, unit)
	} else {
		v.Value[unit] = newQuantity
	}
	return v
}

/**
 * Subtract an array of assets from the Value class's value record. If an asset with the same unit already exists in the value record, the quantity of the
 * existing asset will be decreased by the quantity of the new asset. If no such asset exists, an error message should be printed.
 * @param assets
 * @returns this
 */
func (v *Value) NegateAssets(assets []*rmodels.Asset) *Value {
	for _, asset := range assets {
		v.NegateAsset(asset)
	}
	return v
}

/**
 * Get the quantity of asset object per unit
 * @param unit
 * @returns
 */
func (v *Value) Get(unit string) int64 {
	if value, exists := v.Value[unit]; exists {
		return value
	}
	return 0
}

/**
 * Get all assets (return Record of Asset[])
 * @param
 * @returns Record<string, Asset[]>
 */
func (v *Value) Units() map[string][]*ValueUnit {
	result := make(map[string][]*ValueUnit)
	for unit, quantity := range v.Value {
		if _, exists := result[unit]; !exists {
			result[unit] = []*ValueUnit{}
		}
		result[unit] = append(result[unit], &ValueUnit{
			Unit:     unit,
			Quantity: quantity,
		})
	}
	return result
}

type ValueUnit struct {
	Unit     string
	Quantity int64
}

/**
 * Check if the value is greater than or equal to an inputted value
 * @param unit - The unit to compare (e.g., "ADA")
 * @param other - The value to compare against
 * @returns boolean
 */
// geq = (unit: string, other: Value): boolean => {
//     const thisValue = this.get(unit);
//     const otherValue = other.get(unit);
//     return thisValue >= otherValue;
// };

func (v *Value) Geq(unit string, other *Value) bool {
	if _, exists := v.Value[unit]; !exists {
		return false
	}
	if _, exists := other.Value[unit]; !exists {
		return false
	}
	return v.Value[unit] >= other.Value[unit]
}

/**
 * Check if the value is less than or equal to an inputted value
 * @param unit - The unit to compare (e.g., "ADA")
 * @param other - The value to compare against
 * @returns boolean
 */
// leq = (unit: string, other: Value): boolean => {
//     const thisValue = this.get(unit);
//     const otherValue = other.get(unit);
//     if (otherValue === undefined) {
//         return false;
//     }

//     return thisValue <= otherValue;
// };

func (v *Value) Leq(unit string, other *Value) bool {
	if _, exists := v.Value[unit]; !exists {
		return false
	}
	if _, exists := other.Value[unit]; !exists {
		return false
	}
	return v.Value[unit] <= other.Value[unit]
}

/**
 * Check if the value is empty
 * @param
 * @returns boolean
 */
func (v *Value) IsEmpty() bool {
	return len(v.Value) == 0
}

/**
 * Merge the given values
 * @param values
 * @returns this
 */
func (v *Value) Merge(values ...*Value) *Value {
	for _, other := range values {
		for unit, quantity := range other.Value {
			v.Value[unit] += quantity
		}
	}
	return v
}
