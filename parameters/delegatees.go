package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DelegateesParam interface {
	CalendarUserAddressTypeParameter
}

// NewDelegateesParam create a new registries.DelegatedTo property
// This parameter can be used in this property :
// - registries.Attendee
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.5
func NewDelegateesParam(value *url.URL) DelegateesParam {
	return &calendarUserAddressParameter{
		ParamName: registries.DelegatedTo,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
