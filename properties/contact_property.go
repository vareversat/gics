package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ContactProperty interface {
	TextPropertyType
}

// NewContactProperty create a new registries.ContactProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vfreebusy (Optional)
// Optional parameters :
// - registries.AlternateTextRepresentationParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2
func NewContactProperty(value string, params ...parameters.Parameter) ContactProperty {
	return &textPropertyType{
		PropName:   registries.ContactProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
