package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type OrganizerProperty interface {
}

type organizerProperty struct {
	IANAToken registries.Properties
	Value     values.CalendarUserAddressValue
}

func NewOrganizerProperty(uri url.URL) OrganizerProperty {
	return &organizerProperty{
		IANAToken: registries.ORGANIZER,
		Value:     values.NewCalendarUserAddressValue(uri),
	}
}
