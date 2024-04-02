package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.1

import (
	"io"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type EventCalendarComponent interface {
	CalendarComponent
}

type eventCalendarComponent struct {
	Begin         properties.BlockDelimiterProperty
	UID           properties.UidProperty
	DateTimeStamp properties.DateTimeStampProperty
	Properties    properties.Properties
	End           properties.BlockDelimiterProperty
}

func NewEventCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	propertyList ...properties.Property) EventCalendarComponent {
	return &eventCalendarComponent{
		Begin:         properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VEVENT),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registries.END, properties.VEVENT),
	}
}

func (eC *eventCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	eC.Begin.ToICalendarPropFormat(output)
	eC.UID.ToICalendarPropFormat(output)
	eC.DateTimeStamp.ToICalendarPropFormat(output)
	for i := 0; i < len(eC.Properties); i++ {
		eC.Properties[i].ToICalendarPropFormat(output)
	}
	eC.End.ToICalendarPropFormat(output)
}
