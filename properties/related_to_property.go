package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RelatedToProperty interface {
	TextPropertyType
}

// NewRelatedToProperty create a new registries.RelatedToProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Mandatory if group-scheduled calendar)
// - registries.Vtodo (Mandatory if group-scheduled calendar)
// - registries.Vjournal (Mandatory if group-scheduled calendar)
// Optional parameters :
// - registries.RelationshipTypeParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.5
func NewRelatedToProperty(value string, params ...parameters.Parameter) RelatedToProperty {
	return &textPropertyType{
		PropName:   registries.RelatedToProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
