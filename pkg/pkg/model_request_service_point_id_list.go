// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type RequestServicePointIdList struct {

	Data RequestServicePointIdListData `json:"data"`

	Meta map[string]interface{} `json:"meta,omitempty"`
}

// AssertRequestServicePointIdListRequired checks if the required fields are not zero-ed
func AssertRequestServicePointIdListRequired(obj RequestServicePointIdList) error {
	elements := map[string]interface{}{
		"data": obj.Data,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertRequestServicePointIdListDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertRequestServicePointIdListConstraints checks if the values respects the defined constraints
func AssertRequestServicePointIdListConstraints(obj RequestServicePointIdList) error {
	return nil
}