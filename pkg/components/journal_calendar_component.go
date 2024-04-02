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
	Begin         properties.BlockDelimiterProperty
	UID           properties.UidProperty
	DateTimeStamp properties.DateTimeStampProperty
	Properties    properties.Properties
	End           properties.BlockDelimiterProperty
}

func NewJournalCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	propertyList ...properties.Property) JournalCalendarComponent {
	return &journalCalendarComponent{
		Begin:         properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VJOURNAL),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registries.END, properties.VJOURNAL),
	}
}

func (jC *journalCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	jC.Begin.ToICalendarPropFormat(output)
	jC.UID.ToICalendarPropFormat(output)
	jC.DateTimeStamp.ToICalendarPropFormat(output)
	for i := 0; i < len(jC.Properties); i++ {
		jC.Properties[i].ToICalendarPropFormat(output)
	}
	jC.End.ToICalendarPropFormat(output)
}
