package parameters

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DelegatedFromParam interface {
	CalendarUserAddressTypeParameter
}

// NewDelegatedFromParam create a new registries.DelegatedFromParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.AttendeeProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.4
func NewDelegatedFromParam(value *url.URL) DelegatedFromParam {
	return &calendarUserAddressParameter{
		ParamName: registries.DelegatedFromParam,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}

func NewDelegatedFromParamFromString(value string) (DelegatedFromParam, error) {
	urlValue, err := url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", value)
	}
	return &calendarUserAddressParameter{
		ParamName: registries.DelegatedFromParam,
		Value:     types.NewCalendarUserAddressValue(urlValue),
	}, nil
}
