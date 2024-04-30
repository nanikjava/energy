// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner struct {

	// The type of the charging time period. If absent applies to all periods
	Type string `json:"type,omitempty"`

	// Display name of the tariff
	DisplayName string `json:"displayName"`

	// Array of feed in rates
	Rates []EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner `json:"rates,omitempty"`

	// Usage period for which the block rate applies. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). Defaults to P1Y if absent
	Period string `json:"period,omitempty"`

	// Array of time periods for which this tariff is applicable
	TimeVariations []EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner `json:"timeVariations"`
}

// AssertEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerRequired checks if the required fields are not zero-ed
func AssertEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerRequired(obj EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) error {
	elements := map[string]interface{}{
		"displayName": obj.DisplayName,
		"timeVariations": obj.TimeVariations,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Rates {
		if err := AssertEnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInnerRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TimeVariations {
		if err := AssertEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerConstraints checks if the values respects the defined constraints
func AssertEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerConstraints(obj EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) error {
	return nil
}
