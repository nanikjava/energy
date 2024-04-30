// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyPlanControlledLoadV2Inner struct {

	// A display name for the controlled load
	DisplayName string `json:"displayName"`

	// Specifies the type of controlloed load rate 
	RateBlockUType string `json:"rateBlockUType"`

	// Optional start date of the application of the controlled load rate
	StartDate string `json:"startDate,omitempty"`

	// Optional end date of the application of the controlled load rate
	EndDate string `json:"endDate,omitempty"`

	SingleRate EnergyPlanControlledLoadV2InnerSingleRate `json:"singleRate,omitempty"`

	// Array of objects representing time of use rates.  Required if rateBlockUType is timeOfUseRates
	TimeOfUseRates []EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner `json:"timeOfUseRates,omitempty"`
}

// AssertEnergyPlanControlledLoadV2InnerRequired checks if the required fields are not zero-ed
func AssertEnergyPlanControlledLoadV2InnerRequired(obj EnergyPlanControlledLoadV2Inner) error {
	elements := map[string]interface{}{
		"displayName": obj.DisplayName,
		"rateBlockUType": obj.RateBlockUType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertEnergyPlanControlledLoadV2InnerSingleRateRequired(obj.SingleRate); err != nil {
		return err
	}
	for _, el := range obj.TimeOfUseRates {
		if err := AssertEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEnergyPlanControlledLoadV2InnerConstraints checks if the values respects the defined constraints
func AssertEnergyPlanControlledLoadV2InnerConstraints(obj EnergyPlanControlledLoadV2Inner) error {
	return nil
}
