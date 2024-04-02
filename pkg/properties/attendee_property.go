package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1

import (
	"net/url"

	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type AttendeeProperty interface {
	CalendarUserAddressPropertyType
}

func NewAttendeeProperty(uri *url.URL, params ...parameters.Parameter) AttendeeProperty {
	return &calendarUserAddressPropertyType{
		PropName:   registries.ATTENDEE,
		Value:      types.NewCalendarUserAddressValue(uri),
		Parameters: params,
	}
}
