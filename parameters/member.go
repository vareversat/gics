package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type MemberParam interface {
	CalendarUserAddressTypeParameter
}

// NewMemberParam create a new registries.MemberParam property
// This parameter can be used in this property :
// - registries.Attendee
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.11
func NewMemberParam(value *url.URL) MemberParam {
	return &calendarUserAddressParameter{
		ParamName: registries.MemberParam,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
