// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




// EnergyPlanDiscountsInnerPercentOfBill - Required if methodUType is percentOfBill
type EnergyPlanDiscountsInnerPercentOfBill struct {

	// The rate of the discount applied to the bill amount
	Rate string `json:"rate"`
}

// AssertEnergyPlanDiscountsInnerPercentOfBillRequired checks if the required fields are not zero-ed
func AssertEnergyPlanDiscountsInnerPercentOfBillRequired(obj EnergyPlanDiscountsInnerPercentOfBill) error {
	elements := map[string]interface{}{
		"rate": obj.Rate,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertEnergyPlanDiscountsInnerPercentOfBillConstraints checks if the values respects the defined constraints
func AssertEnergyPlanDiscountsInnerPercentOfBillConstraints(obj EnergyPlanDiscountsInnerPercentOfBill) error {
	return nil
}
