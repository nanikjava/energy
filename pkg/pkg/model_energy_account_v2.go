// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyAccountV2 struct {

	// The ID of the account.  To be created in accordance with CDR ID permanence requirements
	AccountId string `json:"accountId"`

	// Optional identifier of the account as defined by the data holder.  This must be the value presented on physical statements (if it exists) and must not be used for the value of accountId
	AccountNumber string `json:"accountNumber,omitempty"`

	// An optional display name for the account if one exists or can be derived.  The content of this field is at the discretion of the data holder
	DisplayName string `json:"displayName,omitempty"`

	// Open or closed status for the account. If not present then OPEN is assumed
	OpenStatus string `json:"openStatus,omitempty"`

	// The date that the account was created or opened. Mandatory if openStatus is OPEN
	CreationDate string `json:"creationDate,omitempty"`

	// The array of plans containing service points and associated plan details
	Plans []EnergyAccountV2AllOfPlans `json:"plans"`
}

// AssertEnergyAccountV2Required checks if the required fields are not zero-ed
func AssertEnergyAccountV2Required(obj EnergyAccountV2) error {
	elements := map[string]interface{}{
		"accountId": obj.AccountId,
		"plans": obj.Plans,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Plans {
		if err := AssertEnergyAccountV2AllOfPlansRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEnergyAccountV2Constraints checks if the values respects the defined constraints
func AssertEnergyAccountV2Constraints(obj EnergyAccountV2) error {
	return nil
}