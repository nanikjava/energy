// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyServicePointDetailRelatedParticipants struct {

	// The name of the party/organisation related to this service point
	Party string `json:"party"`

	// The role performed by this participant in relation to the service point. Note the details of enumeration values below: <ul><li>**FRMP** - Financially Responsible Market Participant</li><li>**LNSP** - Local Network Service Provider or Embedded Network Manager for child connection points</li><li>**DRSP** - wholesale Demand Response and/or market ancillary Service Provider and note that where it is not relevant for a NMI it will not be included</li></ul>
	Role string `json:"role"`
}

// AssertEnergyServicePointDetailRelatedParticipantsRequired checks if the required fields are not zero-ed
func AssertEnergyServicePointDetailRelatedParticipantsRequired(obj EnergyServicePointDetailRelatedParticipants) error {
	elements := map[string]interface{}{
		"party": obj.Party,
		"role": obj.Role,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertEnergyServicePointDetailRelatedParticipantsConstraints checks if the values respects the defined constraints
func AssertEnergyServicePointDetailRelatedParticipantsConstraints(obj EnergyServicePointDetailRelatedParticipants) error {
	return nil
}