package components

import (
	"io"

	"github.com/vareversat/gics/pkg/types"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

// ToDoCalendarComponent is the interface definition of a VTODO calendar component
// See also : [CalendarComponent]
type ToDoCalendarComponent interface {
	CalendarComponent
}

type toDoCalendarComponent struct {
	Begin         properties.BlockDelimiterProperty
	UID           properties.UidProperty
	DateTimeStamp properties.DateTimeStampProperty
	Properties    properties.Properties
	End           properties.BlockDelimiterProperty
}

// NewToDoCalendarComponent create a VTODO calendar component
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.2
func NewToDoCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	propertyList ...properties.Property) ToDoCalendarComponent {
	return &toDoCalendarComponent{
		Begin:         properties.NewBlockDelimiterProperty(registries.BEGIN, types.VTODO),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registries.END, types.VTODO),
	}
}

func (tC *toDoCalendarComponent) MandatoryProperties() []registries.PropertyNames {
	return []registries.PropertyNames{
		registries.BEGIN,
		registries.END,
		registries.UID,
		registries.DTSTAMP,
	}
}

func (tC *toDoCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{registries.DUE, registries.DURATION_PROP}
}

func (tC *toDoCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{}
}

func (tC *toDoCalendarComponent) SerializeToICSFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	tC.UID.ToICalendarPropFormat(output)
	tC.DateTimeStamp.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}
