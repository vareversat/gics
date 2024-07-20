package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.4

// Parameter used in these properties :
// - ATTENDEE

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DelegatorsParam interface {
	CalendarUserAddressTypeParameter
}

func NewDelegatorsParam(value *url.URL) DelegatorsParam {
	return &calendarUserAddressParameter{
		ParamName: registries.DELEGATEDFROM,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
