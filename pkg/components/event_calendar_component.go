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
	Begin properties.BlockDelimiterProperty
	End   properties.BlockDelimiterProperty

	Properties properties.Properties
}

func NewEventCalendarComponent(propertyList ...properties.Property) EventCalendarComponent {
	return &eventCalendarComponent{
		Begin:      properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VEVENT),
		End:        properties.NewBlockDelimiterProperty(registries.END, properties.VEVENT),
		Properties: propertyList,
	}
}

func (eC *eventCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	eC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(eC.Properties); i++ {
		eC.Properties[i].ToICalendarPropFormat(output)
	}
	eC.End.ToICalendarPropFormat(output)
}
