// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyAccountListResponseV2 struct {

	Data EnergyAccountListResponseV2Data `json:"data"`

	Links LinksPaginated `json:"links"`

	Meta MetaPaginated `json:"meta"`
}

// AssertEnergyAccountListResponseV2Required checks if the required fields are not zero-ed
func AssertEnergyAccountListResponseV2Required(obj EnergyAccountListResponseV2) error {
	elements := map[string]interface{}{
		"data": obj.Data,
		"links": obj.Links,
		"meta": obj.Meta,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertEnergyAccountListResponseV2DataRequired(obj.Data); err != nil {
		return err
	}
	if err := AssertLinksPaginatedRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertMetaPaginatedRequired(obj.Meta); err != nil {
		return err
	}
	return nil
}

// AssertEnergyAccountListResponseV2Constraints checks if the values respects the defined constraints
func AssertEnergyAccountListResponseV2Constraints(obj EnergyAccountListResponseV2) error {
	return nil
}