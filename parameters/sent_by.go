package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.18

// Parameter used in these properties :
// - ATTENDEE
// - ORGANIZER

import (
	"net/url"

	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type SentByParam interface {
	CalendarUserAddressTypeParameter
}

func NewSentByParam(value *url.URL) SentByParam {
	return &calendarUserAddressParameter{
		ParamName: registry.SENTBY,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
