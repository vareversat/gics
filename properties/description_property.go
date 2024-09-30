package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DescriptionProperty interface {
	TextPropertyType
}

// NewDescriptionProperty create a new registries.DescriptionProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional, multiple time)
// - registries.Vfreebusy (Optional)
// Optional parameters :
// - registries.AlternateTextRepresentationParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5
func NewDescriptionProperty(descriptionValue string,
	params ...parameters.Parameter) DescriptionProperty {
	return &textPropertyType{
		PropName:   registries.DescriptionProp,
		Value:      types.NewTextValue(descriptionValue),
		Parameters: params,
	}
}
