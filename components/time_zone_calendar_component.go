package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
)

// TimeZoneCalendarComponent is the interface definition of a VTIMEZONE calendar component
// See also : [Component]
type TimeZoneCalendarComponent interface {
	Component
}

// TimeZoneCalendarSubComponent may be either a STANDARD or a DAYLIGHT subcomponent of a VTIMEZONE calendar component
type TimeZoneCalendarSubComponent interface {
	Component
}

// TimeZoneCalendarSubComponents array of TimeZoneCalendarSubComponent
type TimeZoneCalendarSubComponents []TimeZoneCalendarSubComponent

// TimeZoneCalendarStandardComponent is a STANDARD subcomponent of a VTIMEZONE calendar component
type TimeZoneCalendarStandardComponent interface {
	TimeZoneCalendarSubComponent
}

// TimeZoneCalendarDaylightComponent is a DAYLIGHT subcomponent of a VTIMEZONE calendar component
type TimeZoneCalendarDaylightComponent interface {
	TimeZoneCalendarSubComponent
}

type timeZoneCalendarComponent struct {
	Begin      properties.BeginProperty
	Properties properties.Properties
	Components TimeZoneCalendarSubComponents
	End        properties.EndProperty
}

type timeZoneCalendarSubComponent struct {
	Begin      properties.BeginProperty
	Properties properties.Properties
	End        properties.EndProperty
}

// NewTimeZoneCalendarComponent create a VTIMEZONE calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.5
func NewTimeZoneCalendarComponent(
	components TimeZoneCalendarSubComponents,
	propertyList ...properties.Property) TimeZoneCalendarStandardComponent {
	return &timeZoneCalendarComponent{
		Begin: properties.NewBeginProperty(
			registries.Vtimezone,
		),
		Components: components,
		Properties: propertyList,
		End: properties.NewEndProperty(
			registries.Vtimezone,
		),
	}
}

func (tC *timeZoneCalendarComponent) GetProperty(
	name registries.PropertyRegistry,
) properties.Property {
	for i := 0; i < len(tC.Properties); i++ {
		if tC.Properties[i].GetName() == name {
			return tC.Properties[i]
		}
	}
	return nil
}

func (tC *timeZoneCalendarSubComponent) GetProperty(
	name registries.PropertyRegistry,
) properties.Property {
	for i := 0; i < len(tC.Properties); i++ {
		if tC.Properties[i].GetName() == name {
			return tC.Properties[i]
		}
	}
	return nil
}

// NewTimeZoneDayLightSubcomponent create a STANDARD subcomponent of a VTIMEZONE calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.5
func NewTimeZoneDayLightSubcomponent(
	propertyList ...properties.Property) TimeZoneCalendarDaylightComponent {
	return &timeZoneCalendarSubComponent{
		Begin: properties.NewBeginProperty(
			registries.Daylight,
		),
		Properties: propertyList,
		End: properties.NewEndProperty(
			registries.Daylight,
		),
	}
}

// NewTimeZoneCalendarStandardSubcomponent create a DAYLIGHT subcomponent of a VTIMEZONE calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.5
func NewTimeZoneCalendarStandardSubcomponent(
	propertyList ...properties.Property) TimeZoneCalendarSubComponent {
	return &timeZoneCalendarSubComponent{
		Begin: properties.NewBeginProperty(
			registries.Standard,
		),
		Properties: propertyList,
		End: properties.NewEndProperty(
			registries.Standard,
		),
	}
}

func (tC *timeZoneCalendarComponent) SerializeToICSFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	for i := 0; i < len(tC.Components); i++ {
		tC.Components[i].SerializeToICSFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}

func (tC *timeZoneCalendarComponent) MandatoryProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{
		registries.BeginProp,
		registries.EndProp,
		registries.TimeZoneIdProp,
	}
}

func (tC *timeZoneCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (tC *timeZoneCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (tC *timeZoneCalendarComponent) AddProperty(property properties.Property) {
	tC.Properties = append(tC.Properties, property)
}

func (tC *timeZoneCalendarSubComponent) SerializeToICSFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}

func (tC *timeZoneCalendarSubComponent) MandatoryProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{
		registries.BeginProp,
		registries.EndProp,
		registries.DateTimeStartProp,
		registries.TimeZoneOffsetToProp,
		registries.TimeZoneOffsetFromProp,
	}
}

func (tC *timeZoneCalendarSubComponent) MutuallyExclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (tC *timeZoneCalendarSubComponent) MutuallyInclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (tC *timeZoneCalendarSubComponent) AddProperty(property properties.Property) {
	tC.Properties = append(tC.Properties, property)
}
