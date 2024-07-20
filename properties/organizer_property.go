package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type OrganizerProperty interface {
	CalendarUserAddressPropertyType
}

func NewOrganizerProperty(uri *url.URL, params ...parameters.Parameter) OrganizerProperty {
	return &calendarUserAddressPropertyType{
		PropName:   registry.ORGANIZER,
		Value:      types.NewCalendarUserAddressValue(uri),
		Parameters: params,
	}
}

func NewOrganizerPropertyFromString(
	uri string,
	params ...parameters.Parameter,
) (AttendeeProperty, error) {
	urlValue, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", uri)
	}

	return &calendarUserAddressPropertyType{
		PropName:   registry.ORGANIZER,
		Value:      types.NewCalendarUserAddressValue(urlValue),
		Parameters: params,
	}, nil
}
