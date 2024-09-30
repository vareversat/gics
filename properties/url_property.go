package properties

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type UrlProperty interface {
	UriPropertyType
}

// NewUrlProperty create a new registries.UrlProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vfreebusy (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.6
func NewUrlProperty(uri *url.URL, params ...parameters.Parameter) UrlProperty {
	return &uriPropertyType{
		PropName:   registries.UrlProp,
		Value:      types.NewUriValue(uri),
		Parameters: params,
	}
}

// NewUrlPropertyFromString create a new registries.UrlProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vfreebusy (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.6
func NewUrlPropertyFromString(uri string, params ...parameters.Parameter) (UrlProperty, error) {
	urlValue, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", uri)
	}

	return &uriPropertyType{
		PropName:   registries.UrlProp,
		Value:      types.NewUriValue(urlValue),
		Parameters: params,
	}, nil
}
