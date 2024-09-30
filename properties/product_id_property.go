package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ProductIdProperty interface {
	TextPropertyType
}

// NewProductIdProperty create a new registries.ProductIdentifierProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vcalendar (Mandatory)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.3
func NewProductIdProperty(value string, params ...parameters.Parameter) ProductIdProperty {
	return &textPropertyType{
		PropName:   registries.ProductIdentifierProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
