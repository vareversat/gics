package types

import (
	"net/url"

	"github.com/vareversat/gics/registries"
)

type CalendarUserAddressType interface {
	ValueType
	GetValue() *url.URL
}

type calendarUserAddressType struct {
	typeName  registries.ValueTypeRegistry
	typeValue *url.URL
}

// GetValue get the *url.URL typed value
func (cUAT *calendarUserAddressType) GetValue() *url.URL {
	return cUAT.typeValue
}

func (cUAT *calendarUserAddressType) GetStringValue() string {
	return cUAT.typeValue.String()
}
func (cUAT *calendarUserAddressType) GetTypeName() string {
	return string(cUAT.typeName)
}

// NewCalendarUserAddressValue create a new [registries.CalAddress] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.3
func NewCalendarUserAddressValue(value *url.URL) CalendarUserAddressType {
	return &calendarUserAddressType{
		typeName:  registries.CalAddress,
		typeValue: value,
	}
}
