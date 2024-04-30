// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver


import (
	"fmt"
)



type Effective string

// List of Effective
const (
	CURRENT Effective = "CURRENT"
	FUTURE Effective = "FUTURE"
	ALL_EFFECTIVE Effective = "ALL_EFFECTIVE"
)

// AllowedEffectiveEnumValues is all the allowed values of Effective enum
var AllowedEffectiveEnumValues = []Effective{
	"CURRENT",
	"FUTURE",
	"ALL_EFFECTIVE",
}

// validEffectiveEnumValue provides a map of Effectives for fast verification of use input
var validEffectiveEnumValues = map[Effective]struct{}{
	"CURRENT": {},
	"FUTURE": {},
	"ALL_EFFECTIVE": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Effective) IsValid() bool {
	_, ok := validEffectiveEnumValues[v]
	return ok
}

// NewEffectiveFromValue returns a pointer to a valid Effective
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEffectiveFromValue(v string) (Effective, error) {
	ev := Effective(v)
	if ev.IsValid() {
		return ev, nil
	} else {
		return "", fmt.Errorf("invalid value '%v' for Effective: valid values are %v", v, AllowedEffectiveEnumValues)
	}
}



// AssertEffectiveRequired checks if the required fields are not zero-ed
func AssertEffectiveRequired(obj Effective) error {
	return nil
}

// AssertEffectiveConstraints checks if the values respects the defined constraints
func AssertEffectiveConstraints(obj Effective) error {
	return nil
}
