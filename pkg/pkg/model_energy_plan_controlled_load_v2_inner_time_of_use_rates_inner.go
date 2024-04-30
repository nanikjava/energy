// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner struct {

	// Display name of the controlled load rate
	DisplayName string `json:"displayName"`

	// Description of the controlled load rate
	Description string `json:"description,omitempty"`

	// The daily supply charge (exclusive of GST) for this controlled load tier
	DailySupplyCharge string `json:"dailySupplyCharge,omitempty"`

	// Array of controlled load rates in order of usage volume
	Rates []EnergyPlanControlledLoadV2InnerSingleRateRatesInner `json:"rates"`

	// Usage period for which the block rate applies. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). Defaults to P1Y if absent
	Period string `json:"period,omitempty"`

	// Array of times of use.
	TimeOfUse []EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner `json:"timeOfUse"`

	// The type of usage that the rate applies to
	Type string `json:"type"`
}

// AssertEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerRequired checks if the required fields are not zero-ed
func AssertEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerRequired(obj EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner) error {
	elements := map[string]interface{}{
		"displayName": obj.DisplayName,
		"rates": obj.Rates,
		"timeOfUse": obj.TimeOfUse,
		"type": obj.Type,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Rates {
		if err := AssertEnergyPlanControlledLoadV2InnerSingleRateRatesInnerRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TimeOfUse {
		if err := AssertEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerConstraints checks if the values respects the defined constraints
func AssertEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerConstraints(obj EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner) error {
	return nil
}
