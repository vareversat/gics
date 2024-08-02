package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type SentByParam interface {
	CalendarUserAddressTypeParameter
}

// NewSentByParam create a new registries.SentByParam property
// This parameter can be used in these properties :
// - registries.AttendeeProp
// - registries.OrganizerProp
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.18
func NewSentByParam(value *url.URL) SentByParam {
	return &calendarUserAddressParameter{
		ParamName: registries.SentByParam,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
