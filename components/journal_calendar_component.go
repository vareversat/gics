package components

import (
	"io"

	"github.com/vareversat/gics/types"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
)

// JournalCalendarComponent is the interface definition of a VJOURNAL calendar component
// See also : [CalendarComponent]
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

// NewJournalCalendarComponent create a VJOURNAL calendar component
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.3
func NewJournalCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	propertyList ...properties.Property) JournalCalendarComponent {
	return &journalCalendarComponent{
		Begin:         properties.NewBlockDelimiterProperty(registries.BEGIN, types.VJOURNAL),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registries.END, types.VJOURNAL),
	}
}

func (jC *journalCalendarComponent) GetProperty(name registries.PropertyNames) properties.Property {
	for i := 0; i < len(jC.Properties); i++ {
		if jC.Properties[i].GetName() == name {
			return jC.Properties[i]
		}
	}
	return nil
}

func (jC *journalCalendarComponent) MandatoryProperties() []registries.PropertyNames {
	return []registries.PropertyNames{
		registries.BEGIN,
		registries.END,
		registries.UID,
		registries.DTSTAMP,
	}
}

func (jC *journalCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{}
}

func (jC *journalCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{}
}

func (jC *journalCalendarComponent) SerializeToICSFormat(output io.Writer) {
	jC.Begin.ToICalendarPropFormat(output)
	jC.UID.ToICalendarPropFormat(output)
	jC.DateTimeStamp.ToICalendarPropFormat(output)
	for i := 0; i < len(jC.Properties); i++ {
		jC.Properties[i].ToICalendarPropFormat(output)
	}
	jC.End.ToICalendarPropFormat(output)
}
