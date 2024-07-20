package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type AttendeeProperty interface {
	CalendarUserAddressPropertyType
}

// NewAttendeeProperty create a new ATTENDEE property
// This property CAN be seen in VALARM, VEVENT, VFREEBUSY & VJOURNAL components
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
func NewAttendeeProperty(uri *url.URL, params ...parameters.Parameter) AttendeeProperty {
	return &calendarUserAddressPropertyType{
		PropName:   registry.ATTENDEE,
		Value:      types.NewCalendarUserAddressValue(uri),
		Parameters: params,
	}
}

func NewAttendeePropertyFromString(
	uri string,
	params ...parameters.Parameter,
) (OrganizerProperty, error) {
	urlValue, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", uri)
	}

	return &calendarUserAddressPropertyType{
		PropName:   registry.ATTENDEE,
		Value:      types.NewCalendarUserAddressValue(urlValue),
		Parameters: params,
	}, nil
}
