// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type ResponseErrorListV2Errors struct {

	// The code of the error encountered. Where the error is specific to the respondent, an application-specific error code, expressed as a string value. If the error is application-specific, the URN code that the specific error extends must be provided in the meta object. Otherwise, the value is the error code URN.
	Code string `json:"code"`

	// A short, human-readable summary of the problem that MUST NOT change from occurrence to occurrence of the problem represented by the error code.
	Title string `json:"title"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`

	Meta ResponseErrorListV2Meta `json:"meta,omitempty"`
}

// AssertResponseErrorListV2ErrorsRequired checks if the required fields are not zero-ed
func AssertResponseErrorListV2ErrorsRequired(obj ResponseErrorListV2Errors) error {
	elements := map[string]interface{}{
		"code": obj.Code,
		"title": obj.Title,
		"detail": obj.Detail,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertResponseErrorListV2MetaRequired(obj.Meta); err != nil {
		return err
	}
	return nil
}

// AssertResponseErrorListV2ErrorsConstraints checks if the values respects the defined constraints
func AssertResponseErrorListV2ErrorsConstraints(obj ResponseErrorListV2Errors) error {
	return nil
}