package properties

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type BeginProperty interface {
	BlockDelimiterPropertyType
}

// NewBeginProperty create a new registries.BeginProp property.
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
func NewBeginProperty(registry registries.ComponentRegistry) BeginProperty {
	return &blockDelimiterPropertyType{
		PropName: registries.BeginProp,
		Value:    types.NewBlockDelimiterValue(registry),
	}
}

// NewBeginPropertyFromString create a new registries.BeginProp property.
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
func NewBeginPropertyFromString(registry string) BeginProperty {
	return &blockDelimiterPropertyType{
		PropName: registries.BeginProp,
		Value:    types.NewBlockDelimiterValue(registries.ComponentRegistry(registry)),
	}
}
