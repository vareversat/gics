package parameters

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DelegatedToParam interface {
	CalendarUserAddressTypeParameter
}

// NewDelegatedToParam create a new registries.DelegatedToParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.AttendeeProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.5
func NewDelegatedToParam(value *url.URL) DelegatedToParam {
	return &calendarUserAddressParameter{
		ParamName: registries.DelegatedToParam,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}

func NewDelegatedToParamFromString(value string) (DelegatedToParam, error) {
	urlValue, err := url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", value)
	}
	return &calendarUserAddressParameter{
		ParamName: registries.DelegatedToParam,
		Value:     types.NewCalendarUserAddressValue(urlValue),
	}, nil
}
