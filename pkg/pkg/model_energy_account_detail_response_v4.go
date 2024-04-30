// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyAccountDetailResponseV4 struct {

	Data EnergyAccountDetailV4 `json:"data"`

	Links Links `json:"links"`

	Meta map[string]interface{} `json:"meta,omitempty"`
}

// AssertEnergyAccountDetailResponseV4Required checks if the required fields are not zero-ed
func AssertEnergyAccountDetailResponseV4Required(obj EnergyAccountDetailResponseV4) error {
	elements := map[string]interface{}{
		"data": obj.Data,
		"links": obj.Links,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertEnergyAccountDetailV4Required(obj.Data); err != nil {
		return err
	}
	if err := AssertLinksRequired(obj.Links); err != nil {
		return err
	}
	return nil
}

// AssertEnergyAccountDetailResponseV4Constraints checks if the values respects the defined constraints
func AssertEnergyAccountDetailResponseV4Constraints(obj EnergyAccountDetailResponseV4) error {
	return nil
}
