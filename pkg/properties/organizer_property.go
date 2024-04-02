package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3

import (
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
