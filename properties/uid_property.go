package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type UidProperty interface {
	TextPropertyType
}

// NewUidProperty create a new registries.UidProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Mandatory)
// - registries.Vtodo (Mandatory)
// - registries.Vjournal (Mandatory)
// - registries.Vfreebusy (Mandatory)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7
func NewUidProperty(value string, params ...parameters.Parameter) UidProperty {
	return &textPropertyType{
		PropName:   registries.UidProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
