package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.2

import (
	"io"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type ToDoCalendarComponent interface {
	CalendarComponent
}

type toDoCalendarComponent struct {
	Begin properties.BlockDelimiterProperty
	End   properties.BlockDelimiterProperty

	Properties properties.Properties
}

func NewToDoCalendarComponent(propertyList ...properties.Property) ToDoCalendarComponent {
	return &toDoCalendarComponent{
		Begin:      properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VTODO),
		End:        properties.NewBlockDelimiterProperty(registries.END, properties.VTODO),
		Properties: propertyList,
	}
}

func (tC *toDoCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}
