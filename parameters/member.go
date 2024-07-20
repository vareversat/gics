package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.11

// Parameter used in these properties :
// - ATTENDEE

import (
	"net/url"

	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type MemberParam interface {
	CalendarUserAddressTypeParameter
}

func NewMemberParam(value *url.URL) MemberParam {
	return &calendarUserAddressParameter{
		ParamName: registry.MEMBER,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}
