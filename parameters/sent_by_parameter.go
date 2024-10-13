package parameters

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type SentByParam interface {
	CalendarUserAddressTypeParameter
}

// NewSentByParam create a new registries.SentByParam property. See [RFC-5545] ref for more info
// This parameter can be used in these properties :
// - registries.AttendeeProp
// - registries.OrganizerProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.18
func NewSentByParam(value *url.URL) SentByParam {
	return &calendarUserAddressParameter{
		ParamName: registries.SentByParam,
		Value:     types.NewCalendarUserAddressValue(value),
	}
}

func NewSentByParamFromString(value string) (SentByParam, error) {
	urlValue, err := url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", value)
	}
	return &calendarUserAddressParameter{
		ParamName: registries.SentByParam,
		Value:     types.NewCalendarUserAddressValue(urlValue),
	}, nil
}
