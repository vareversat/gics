package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.3

// Example :mailto:jane_doe@example.com

import (
	"net/url"

	"github.com/vareversat/gics/registry"
)

type CalendarUserAddressValue struct {
	V
	Value *url.URL
}

func NewCalendarUserAddressValue(value *url.URL) CalendarUserAddressValue {
	return CalendarUserAddressValue{
		V:     NewValue(registry.CALADDRESS),
		Value: value,
	}
}

func (cUAV *CalendarUserAddressValue) GetValue() string {
	return cUAV.Value.String()
}
