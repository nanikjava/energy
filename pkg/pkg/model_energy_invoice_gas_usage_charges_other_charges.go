// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyInvoiceGasUsageChargesOtherCharges struct {

	// Type of charge. Assumed to be other if absent
	Type string `json:"type,omitempty"`

	// The aggregate total of charges for this item (exclusive of GST)
	Amount string `json:"amount"`

	// A free text description of the type of charge
	Description string `json:"description"`
}

// AssertEnergyInvoiceGasUsageChargesOtherChargesRequired checks if the required fields are not zero-ed
func AssertEnergyInvoiceGasUsageChargesOtherChargesRequired(obj EnergyInvoiceGasUsageChargesOtherCharges) error {
	elements := map[string]interface{}{
		"amount": obj.Amount,
		"description": obj.Description,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertEnergyInvoiceGasUsageChargesOtherChargesConstraints checks if the values respects the defined constraints
func AssertEnergyInvoiceGasUsageChargesOtherChargesConstraints(obj EnergyInvoiceGasUsageChargesOtherCharges) error {
	return nil
}