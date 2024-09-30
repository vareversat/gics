package properties

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneUrlProperty interface {
	UriPropertyType
}

// NewTimeZoneUrlProperty create a new registries.TimeZoneUrlProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtimezone (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.5
func NewTimeZoneUrlProperty(value *url.URL, params ...parameters.Parameter) TimeZoneUrlProperty {
	return &uriPropertyType{
		PropName:   registries.TimeZoneUrlProp,
		Value:      types.NewUriValue(value),
		Parameters: params,
	}
}

// NewTimeZoneUrlPropertyFromString create a new registries.TimeZoneUrlProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtimezone (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.5
func NewTimeZoneUrlPropertyFromString(
	uri string,
	params ...parameters.Parameter,
) (TimeZoneUrlProperty, error) {
	urlValue, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", uri)
	}

	return &uriPropertyType{
		PropName:   registries.TimeZoneUrlProp,
		Value:      types.NewUriValue(urlValue),
		Parameters: params,
	}, nil
}
