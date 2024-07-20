package components

import (
	"io"

	"github.com/vareversat/gics/types"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registry"
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
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.3
func NewJournalCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	propertyList ...properties.Property) JournalCalendarComponent {
	return &journalCalendarComponent{
		Begin:         properties.NewBlockDelimiterProperty(registry.BEGIN, types.VJOURNAL),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registry.END, types.VJOURNAL),
	}
}

func (jC *journalCalendarComponent) GetProperty(name registry.PropertyNames) properties.Property {
	for i := 0; i < len(jC.Properties); i++ {
		if jC.Properties[i].GetName() == name {
			return jC.Properties[i]
		}
	}
	return nil
}

func (jC *journalCalendarComponent) MandatoryProperties() []registry.PropertyNames {
	return []registry.PropertyNames{
		registry.BEGIN,
		registry.END,
		registry.UID,
		registry.DTSTAMP,
	}
}

func (jC *journalCalendarComponent) MutuallyExclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}

func (jC *journalCalendarComponent) MutuallyInclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
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
