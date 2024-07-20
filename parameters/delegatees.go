package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.5

// Parameter used in these properties :
// - ATTENDEE

import (
	"net/url"

	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type DelegateesParam interface {
	CalendarUserAddressTypeParameter
}

func NewDelegateesParam(value *url.URL) DelegateesParam {
	return &calendarUserAddressParameter{
		ParamName: registry.DelegatedTo,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
