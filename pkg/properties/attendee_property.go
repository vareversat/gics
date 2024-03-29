package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type AttendeeProperty interface{}

type attendeeProperty struct {
	IANAToken registries.Properties
	Value     values.CalendarUserAddressValue
}

func NewAttendeeProperty(uri url.URL) AttendeeProperty {
	return &attendeeProperty{
		IANAToken: registries.ATTENDEE,
		Value:     values.NewCalendarUserAddressValue(uri),
	}
}
