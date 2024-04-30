// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




// CommonSimpleAddress - Required if addressUType is set to simple
type CommonSimpleAddress struct {

	// Name of the individual or business formatted for inclusion in an address used for physical mail
	MailingName string `json:"mailingName,omitempty"`

	// First line of the standard address object
	AddressLine1 string `json:"addressLine1"`

	// Second line of the standard address object
	AddressLine2 string `json:"addressLine2,omitempty"`

	// Third line of the standard address object
	AddressLine3 string `json:"addressLine3,omitempty"`

	// Mandatory for Australian addresses
	Postcode string `json:"postcode,omitempty"`

	// Name of the city or locality
	City string `json:"city"`

	// Free text if the country is not Australia. If country is Australia then must be one of the values defined by the [State Type Abbreviation](https://auspost.com.au/content/dam/auspost_corp/media/documents/australia-post-data-guide.pdf) in the PAF file format. NSW, QLD, VIC, NT, WA, SA, TAS, ACT, AAT
	State string `json:"state"`

	// A valid [ISO 3166 Alpha-3](https://www.iso.org/iso-3166-country-codes.html) country code. Australia (AUS) is assumed if country is not present.
	Country string `json:"country,omitempty"`
}

// AssertCommonSimpleAddressRequired checks if the required fields are not zero-ed
func AssertCommonSimpleAddressRequired(obj CommonSimpleAddress) error {
	elements := map[string]interface{}{
		"addressLine1": obj.AddressLine1,
		"city": obj.City,
		"state": obj.State,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCommonSimpleAddressConstraints checks if the values respects the defined constraints
func AssertCommonSimpleAddressConstraints(obj CommonSimpleAddress) error {
	return nil
}
