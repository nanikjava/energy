// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




// EnergyAccountV2AllOfPlanOverview - Mandatory if openStatus is OPEN
type EnergyAccountV2AllOfPlanOverview struct {

	// The name of the plan if one exists
	DisplayName string `json:"displayName,omitempty"`

	// The start date of the applicability of this plan
	StartDate string `json:"startDate"`

	// The end date of the applicability of this plan
	EndDate string `json:"endDate,omitempty"`
}

// AssertEnergyAccountV2AllOfPlanOverviewRequired checks if the required fields are not zero-ed
func AssertEnergyAccountV2AllOfPlanOverviewRequired(obj EnergyAccountV2AllOfPlanOverview) error {
	elements := map[string]interface{}{
		"startDate": obj.StartDate,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertEnergyAccountV2AllOfPlanOverviewConstraints checks if the values respects the defined constraints
func AssertEnergyAccountV2AllOfPlanOverviewConstraints(obj EnergyAccountV2AllOfPlanOverview) error {
	return nil
}
