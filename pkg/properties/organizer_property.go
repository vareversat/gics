package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type OrganizerProperty interface {
	CalendarUserAddressPropertyType
}

func NewOrganizerProperty(uri *url.URL, params ...parameters.Parameter) OrganizerProperty {
	return &calendarUserAddressPropertyType{
		PropName:   registries.ORGANIZER,
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
		PropName:   registries.ORGANIZER,
		Value:      types.NewCalendarUserAddressValue(urlValue),
		Parameters: params,
	}, nil
}
