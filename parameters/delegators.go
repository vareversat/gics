package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DelegatorsParam interface {
	CalendarUserAddressTypeParameter
}

// NewDelegatorsParam create a new registries.DelegatedFromParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.AttendeeProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.4
func NewDelegatorsParam(value *url.URL) DelegatorsParam {
	return &calendarUserAddressParameter{
		ParamName: registries.DelegatedFromParam,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
