package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.5

// Parameter used in these properties :
// - ATTENDEE

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type DelegateesParam interface {
	CalendarUserAddressParameter
}

func NewDelegateesParam(value *url.URL) DelegateesParam {
	return &calendarUserAddressParameter{
		ParamName: registries.DELEGATEDTO,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
