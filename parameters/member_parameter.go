package parameters

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type MemberParam interface {
	CalendarUserAddressTypeParameter
}

// NewMemberParam create a new registries.MemberParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.AttendeeProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.11
func NewMemberParam(value *url.URL) MemberParam {
	return &calendarUserAddressParameter{
		ParamName: registries.MemberParam,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}

func NewMemberParamFromString(value string) (MemberParam, error) {
	urlValue, err := url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", value)
	}
	return &calendarUserAddressParameter{
		ParamName: registries.MemberParam,
		Value:     types.NewCalendarUserAddressValue(urlValue),
	}, nil
}
