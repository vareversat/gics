package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.3

// Example :mailto:jane_doe@example.com

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
)

type CalendarUserAddressValue struct {
	V
	Value *url.URL
}

func NewCalendarUserAddressValue(value *url.URL) CalendarUserAddressValue {
	return CalendarUserAddressValue{
		V:     NewValue(registries.CALADDRESS),
		Value: value,
	}
}

func (cUAV *CalendarUserAddressValue) GetValue() string {
	return cUAV.Value.String()
}
