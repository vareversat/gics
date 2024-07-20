package components

import (
	"io"

	"github.com/vareversat/gics/types"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registry"
)

// TimeZoneCalendarComponent is the interface definition of a VTIMEZONE calendar component
// See also : [CalendarComponent]
type TimeZoneCalendarComponent interface {
	CalendarComponent
}

// TimeZoneCalendarSubComponent may be either a STANDARD or a DAYLIGHT subcomponent of a VTIMEZONE calendar component
type TimeZoneCalendarSubComponent interface {
	CalendarComponent
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
	Begin      properties.BlockDelimiterProperty
	TimeZoneId properties.TimeZoneIdProperty
	Properties properties.Properties
	Components TimeZoneCalendarSubComponents
	End        properties.BlockDelimiterProperty
}

type timeZoneCalendarSubComponent struct {
	Begin              properties.BlockDelimiterProperty
	DateTimeStart      properties.DateTimeStartProperty
	TimeZoneOffsetTo   properties.TimeZoneOffsetToProperty
	TimeZoneOffsetFrom properties.TimeZoneOffsetFromProperty
	Properties         properties.Properties
	End                properties.BlockDelimiterProperty
}

// NewTimeZoneCalendarComponent create a VTIMEZONE calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.5
func NewTimeZoneCalendarComponent(
	timeZoneId properties.TimeZoneIdProperty,
	components TimeZoneCalendarSubComponents,
	propertyList ...properties.Property) TimeZoneCalendarStandardComponent {
	return &timeZoneCalendarComponent{
		Begin:      properties.NewBlockDelimiterProperty(registry.BEGIN, types.VTIMEZONE),
		TimeZoneId: timeZoneId,
		Components: components,
		Properties: propertyList,
		End:        properties.NewBlockDelimiterProperty(registry.END, types.VTIMEZONE),
	}
}

func (tC *timeZoneCalendarComponent) GetProperty(
	name registry.PropertyNames,
) properties.Property {
	for i := 0; i < len(tC.Properties); i++ {
		if tC.Properties[i].GetName() == name {
			return tC.Properties[i]
		}
	}
	return nil
}

func (tC *timeZoneCalendarSubComponent) GetProperty(
	name registry.PropertyNames,
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
	dateTimeStart properties.DateTimeStartProperty,
	timeZoneOffsetTo properties.TimeZoneOffsetToProperty,
	timeZoneOffsetFrom properties.TimeZoneOffsetFromProperty,
	propertyList ...properties.Property) TimeZoneCalendarDaylightComponent {
	return &timeZoneCalendarSubComponent{
		Begin: properties.NewBlockDelimiterProperty(
			registry.BEGIN,
			types.DAYLIGHT,
		),
		DateTimeStart:      dateTimeStart,
		TimeZoneOffsetTo:   timeZoneOffsetTo,
		TimeZoneOffsetFrom: timeZoneOffsetFrom,
		Properties:         propertyList,
		End: properties.NewBlockDelimiterProperty(
			registry.END,
			types.DAYLIGHT,
		),
	}
}

// NewTimeZoneCalendarStandardSubcomponent create a DAYLIGHT subcomponent of a VTIMEZONE calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.5
func NewTimeZoneCalendarStandardSubcomponent(
	dateTimeStart properties.DateTimeStartProperty,
	timeZoneOffsetTo properties.TimeZoneOffsetToProperty,
	timeZoneOffsetFrom properties.TimeZoneOffsetFromProperty,
	propertyList ...properties.Property) TimeZoneCalendarSubComponent {
	return &timeZoneCalendarSubComponent{
		Begin: properties.NewBlockDelimiterProperty(
			registry.BEGIN,
			types.STANDARD,
		),
		DateTimeStart:      dateTimeStart,
		TimeZoneOffsetTo:   timeZoneOffsetTo,
		TimeZoneOffsetFrom: timeZoneOffsetFrom,
		Properties:         propertyList,
		End: properties.NewBlockDelimiterProperty(
			registry.END,
			types.STANDARD,
		),
	}
}

func (tC *timeZoneCalendarComponent) SerializeToICSFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	tC.TimeZoneId.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.Components); i++ {
		tC.Components[i].SerializeToICSFormat(output)
	}
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}

func (tC *timeZoneCalendarComponent) MandatoryProperties() []registry.PropertyNames {
	return []registry.PropertyNames{registry.BEGIN, registry.END, registry.PROP_TZID}
}

func (tC *timeZoneCalendarComponent) MutuallyExclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}

func (tC *timeZoneCalendarComponent) MutuallyInclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}

func (tC *timeZoneCalendarSubComponent) SerializeToICSFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	tC.DateTimeStart.ToICalendarPropFormat(output)
	tC.TimeZoneOffsetTo.ToICalendarPropFormat(output)
	tC.TimeZoneOffsetFrom.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}

func (tC *timeZoneCalendarSubComponent) MandatoryProperties() []registry.PropertyNames {
	return []registry.PropertyNames{
		registry.BEGIN,
		registry.END,
		registry.DTSTART,
		registry.TZOFFSETTO,
		registry.TZOFFSETFROM,
	}
}

func (tC *timeZoneCalendarSubComponent) MutuallyExclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}

func (tC *timeZoneCalendarSubComponent) MutuallyInclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}
