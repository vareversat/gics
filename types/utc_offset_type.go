package types

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/registries"
)

type UtcOffsetType interface {
	ValueType
	GetValue() int
}

type utcOffsetType struct {
	typeName registries.ValueTypeRegistry
	// Offset in seconds
	typeValue int
}

func (u *utcOffsetType) GetValue() int {
	return u.typeValue
}

func (u *utcOffsetType) GetStringValue() string {
	// Create a fake zone based on the offset in seconds
	loc := time.FixedZone("FIXED_ZONE", u.typeValue)
	t := time.Date(1970, 01, 01, 00, 00, 00, 00, loc)
	return t.Format("-0700")
}

// GetTypeName implements UtcOffsetType.
func (u *utcOffsetType) GetTypeName() string {
	return string(u.typeName)
}

// parseStringToDate take a string value and return an int (offset representation in seconds)
func parseStringToUtcOffset(value string) int {
	time, err := time.Parse("-0700", value)
	if err != nil {
		fmt.Println(err)
	}
	_, offset := time.Zone()
	return offset
}

// NewTextValue create a new [registries.UTCOffset] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.14
func NewUtcOffsetValue(value int) UtcOffsetType {
	return &utcOffsetType{
		typeName:  registries.UTCOffset,
		typeValue: value,
	}
}

// NewTextValue create a new [registries.UTCOffset] type value from string. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.14
func NewUtcOffsetValueFromString(value string) UtcOffsetType {
	return &utcOffsetType{
		typeName:  registries.UTCOffset,
		typeValue: parseStringToUtcOffset(value),
	}
}
