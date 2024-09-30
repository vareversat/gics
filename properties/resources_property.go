package properties

import (
	"strings"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ResourcesProperty interface {
	TextPropertyType
}

// NewResourcesProperty create a new registries.ResourcesProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// Optional parameters :
// - registries.AlternateTextRepresentationParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.10
func NewResourcesProperty(values []string, params ...parameters.Parameter) ResourcesProperty {
	return &textPropertyType{
		PropName:   registries.ResourcesProp,
		Values:     types.NewTextValues(values),
		Parameters: params,
	}
}

// NewResourcesPropertyFromString create a new registries.ResourcesProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// Optional parameters :
// - registries.AlternateTextRepresentationParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.10
func NewResourcesPropertyFromString(
	values string,
	params ...parameters.Parameter,
) ResourcesProperty {
	resources := strings.Split(values, ",")
	return &textPropertyType{
		PropName:   registries.ResourcesProp,
		Values:     types.NewTextValues(resources),
		Parameters: params,
	}
}
