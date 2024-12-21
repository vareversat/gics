package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
)

// JournalCalendarComponent is the interface definition of a VJOURNAL calendar component
// See also : [CalendarComponent]
type JournalCalendarComponent interface {
	CalendarComponent
}

type journalCalendarComponent struct {
	Begin      properties.BeginProperty
	Properties properties.Properties
	End        properties.EndProperty
}

// NewJournalCalendarComponent create a VJOURNAL calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.3
func NewJournalCalendarComponent(
	propertyList ...properties.Property) JournalCalendarComponent {
	return &journalCalendarComponent{
		Begin: properties.NewBeginProperty(
			registries.Vjournal,
		),
		Properties: propertyList,
		End: properties.NewEndProperty(
			registries.Vjournal,
		),
	}
}

func (jC *journalCalendarComponent) GetProperty(
	name registries.PropertyRegistry,
) properties.Property {
	for i := 0; i < len(jC.Properties); i++ {
		if jC.Properties[i].GetName() == name {
			return jC.Properties[i]
		}
	}
	return nil
}

func (jC *journalCalendarComponent) MandatoryProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{
		registries.BeginProp,
		registries.EndProp,
		registries.UidProp,
		registries.DateTimeStampProp,
	}
}

func (jC *journalCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (jC *journalCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (jC *journalCalendarComponent) SerializeToICSFormat(output io.Writer) {
	jC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(jC.Properties); i++ {
		jC.Properties[i].ToICalendarPropFormat(output)
	}
	jC.End.ToICalendarPropFormat(output)
}
