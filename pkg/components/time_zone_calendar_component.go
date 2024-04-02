package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6

import (
	"io"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type TimeZoneCalendarComponent interface {
	CalendarComponent
}

type TimeZoneCalendarSubComponent interface {
	CalendarComponent
}

type TimeZoneCalendarSubComponents []TimeZoneCalendarSubComponent

type TimeZoneCalendarStandardComponent interface {
	TimeZoneCalendarSubComponent
}

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

func NewTimeZoneCalendarComponent(
	timeZoneId properties.TimeZoneIdProperty,
	components TimeZoneCalendarSubComponents,
	propertyList ...properties.Property) TimeZoneCalendarStandardComponent {
	return &timeZoneCalendarComponent{
		Begin:      properties.NewBlockDelimiterProperty(registries.BEGIN, properties.STANDARD),
		TimeZoneId: timeZoneId,
		Components: components,
		Properties: propertyList,
		End:        properties.NewBlockDelimiterProperty(registries.END, properties.STANDARD),
	}
}

func NewTimeZoneDayLightComponent(
	dateTimeStart properties.DateTimeStartProperty,
	timeZoneOffsetTo properties.TimeZoneOffsetToProperty,
	timeZoneOffsetFrom properties.TimeZoneOffsetFromProperty,
	propertyList ...properties.Property) TimeZoneCalendarDaylightComponent {
	return &timeZoneCalendarSubComponent{
		Begin: properties.NewBlockDelimiterProperty(
			registries.BEGIN,
			properties.DAYLIGHT,
		),
		DateTimeStart:      dateTimeStart,
		TimeZoneOffsetTo:   timeZoneOffsetTo,
		TimeZoneOffsetFrom: timeZoneOffsetFrom,
		Properties:         propertyList,
		End: properties.NewBlockDelimiterProperty(
			registries.END,
			properties.DAYLIGHT,
		),
	}
}

func NewTimeZoneCalendarStandardComponent(
	dateTimeStart properties.DateTimeStartProperty,
	timeZoneOffsetTo properties.TimeZoneOffsetToProperty,
	timeZoneOffsetFrom properties.TimeZoneOffsetFromProperty,
	propertyList ...properties.Property) TimeZoneCalendarComponent {
	return &timeZoneCalendarSubComponent{
		Begin: properties.NewBlockDelimiterProperty(
			registries.BEGIN,
			properties.VTIMEZONE,
		),
		DateTimeStart:      dateTimeStart,
		TimeZoneOffsetTo:   timeZoneOffsetTo,
		TimeZoneOffsetFrom: timeZoneOffsetFrom,
		Properties:         propertyList,
		End: properties.NewBlockDelimiterProperty(
			registries.END,
			properties.VTIMEZONE,
		),
	}
}

func (tC *timeZoneCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	tC.TimeZoneId.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	for i := 0; i < len(tC.Components); i++ {
		tC.Components[i].ToICalendarComponentFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}

func (tC *timeZoneCalendarSubComponent) ToICalendarComponentFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	tC.DateTimeStart.ToICalendarPropFormat(output)
	tC.TimeZoneOffsetTo.ToICalendarPropFormat(output)
	tC.TimeZoneOffsetFrom.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}
