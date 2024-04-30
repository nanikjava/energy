// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyUsageListResponseData struct {

	// Array of meter reads sorted by NMI in ascending order followed by readStartDate in descending order
	Reads []EnergyUsageRead `json:"reads"`
}

// AssertEnergyUsageListResponseDataRequired checks if the required fields are not zero-ed
func AssertEnergyUsageListResponseDataRequired(obj EnergyUsageListResponseData) error {
	elements := map[string]interface{}{
		"reads": obj.Reads,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Reads {
		if err := AssertEnergyUsageReadRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEnergyUsageListResponseDataConstraints checks if the values respects the defined constraints
func AssertEnergyUsageListResponseDataConstraints(obj EnergyUsageListResponseData) error {
	return nil
}