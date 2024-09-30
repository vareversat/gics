package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type SummaryProperty interface {
	TextPropertyType
}

// NewSummaryProperty create a new registries.SummaryProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// Optional parameters :
// - registries.AlternateTextRepresentationParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12
func NewSummaryProperty(
	summaryValue string, params ...parameters.Parameter,
) SummaryProperty {
	return &textPropertyType{
		PropName:   registries.SummaryProp,
		Value:      types.NewTextValue(summaryValue),
		Parameters: params,
	}
}
