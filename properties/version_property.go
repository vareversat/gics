package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type VersionProperty interface {
	TextPropertyType
}

// NewVersionProperty create a new registries.VersionProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vcalendar (Mandatory)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4
func NewVersionProperty(value string, params ...parameters.Parameter) VersionProperty {
	return &textPropertyType{
		PropName:   registries.VersionProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
