package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.3

import (
	"io"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type JournalCalendarComponent interface {
	CalendarComponent
}

type journalCalendarComponent struct {
	Begin properties.BlockDelimiterProperty
	End   properties.BlockDelimiterProperty

	Properties properties.Properties
}

func NewJournalCalendarComponent(propertyList ...properties.Property) JournalCalendarComponent {
	return &journalCalendarComponent{
		Begin:      properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VJOURNAL),
		End:        properties.NewBlockDelimiterProperty(registries.END, properties.VJOURNAL),
		Properties: propertyList,
	}
}

func (jC *journalCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	jC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(jC.Properties); i++ {
		jC.Properties[i].ToICalendarPropFormat(output)
	}
	jC.End.ToICalendarPropFormat(output)
}
