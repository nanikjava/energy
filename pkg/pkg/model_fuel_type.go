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



type FuelType string

// List of FuelType
const (
	ELECTRICITY FuelType = "ELECTRICITY"
	GAS FuelType = "GAS"
	DUAL FuelType = "DUAL"
	ALL_FUEL FuelType = "ALL_FUEL"
)

// AllowedFuelTypeEnumValues is all the allowed values of FuelType enum
var AllowedFuelTypeEnumValues = []FuelType{
	"ELECTRICITY",
	"GAS",
	"DUAL",
	"ALL_FUEL",
}

// validFuelTypeEnumValue provides a map of FuelTypes for fast verification of use input
var validFuelTypeEnumValues = map[FuelType]struct{}{
	"ELECTRICITY": {},
	"GAS": {},
	"DUAL": {},
	"ALL_FUEL": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FuelType) IsValid() bool {
	_, ok := validFuelTypeEnumValues[v]
	return ok
}

// NewFuelTypeFromValue returns a pointer to a valid FuelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFuelTypeFromValue(v string) (FuelType, error) {
	ev := FuelType(v)
	if ev.IsValid() {
		return ev, nil
	} else {
		return "", fmt.Errorf("invalid value '%v' for FuelType: valid values are %v", v, AllowedFuelTypeEnumValues)
	}
}



// AssertFuelTypeRequired checks if the required fields are not zero-ed
func AssertFuelTypeRequired(obj FuelType) error {
	return nil
}

// AssertFuelTypeConstraints checks if the values respects the defined constraints
func AssertFuelTypeConstraints(obj FuelType) error {
	return nil
}
