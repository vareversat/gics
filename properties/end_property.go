package properties

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type EndProperty interface {
	BlockDelimiterPropertyType
}

// NewEndProperty create a new registries.EndProp property.
// Usage :
// - registries.Valarm (Mandatory)
// - registries.Vcalendar (Mandatory)
// - registries.Vevent (Mandatory)
// - registries.Vtodo (Mandatory)
// - registries.Vfreebusy (Mandatory)
// - registries.Valarm (Mandatory)
// - registries.Vjournal (Mandatory)
// - registries.Standard (Mandatory)
// - registries.Daylight (Mandatory)
func NewEndProperty(registry registries.ComponentRegistry) EndProperty {
	return &blockDelimiterPropertyType{
		PropName: registries.EndProp,
		Value:    types.NewComponentDelimiterValue(registry),
	}
}

// NewEndPropertyFromString create a new registries.EndProp property.
// Usage :
// - registries.Valarm (Mandatory)
// - registries.Vcalendar (Mandatory)
// - registries.Vevent (Mandatory)
// - registries.Vtodo (Mandatory)
// - registries.Vfreebusy (Mandatory)
// - registries.Valarm (Mandatory)
// - registries.Vjournal (Mandatory)
// - registries.Standard (Mandatory)
// - registries.Daylight (Mandatory)
func NewEndPropertyFromString(registry string) EndProperty {
	return &blockDelimiterPropertyType{
		PropName: registries.EndProp,
		Value:    types.NewComponentDelimiterValue(registries.ComponentRegistry(registry)),
	}
}
