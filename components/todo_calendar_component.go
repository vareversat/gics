package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registry"
)

// ToDoCalendarComponent is the interface definition of a VTODO calendar component
// See also : [CalendarComponent]
type ToDoCalendarComponent interface {
	CalendarComponent
}

type toDoCalendarComponent struct {
	Begin                   properties.BlockDelimiterProperty
	UID                     properties.UidProperty
	DateTimeStamp           properties.DateTimeStampProperty
	AlarmCalendarComponents []AlarmCalendarComponent
	Properties              properties.Properties
	End                     properties.BlockDelimiterProperty
}

// NewToDoCalendarComponent create a VTODO calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.2
func NewToDoCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	alarmCalendarComponents []AlarmCalendarComponent,
	propertyList ...properties.Property) ToDoCalendarComponent {
	return &toDoCalendarComponent{
		Begin: properties.NewBlockDelimiterProperty(
			registry.BEGIN,
			registry.Vtodo,
		),
		UID:                     uid,
		DateTimeStamp:           dateTimeStamp,
		AlarmCalendarComponents: alarmCalendarComponents,
		Properties:              propertyList,
		End:                     properties.NewBlockDelimiterProperty(registry.END, registry.Vtodo),
	}
}

func (tC *toDoCalendarComponent) GetProperty(name registry.PropertyNames) properties.Property {
	for i := 0; i < len(tC.Properties); i++ {
		if tC.Properties[i].GetName() == name {
			return tC.Properties[i]
		}
	}
	return nil
}

func (tC *toDoCalendarComponent) MandatoryProperties() []registry.PropertyNames {
	return []registry.PropertyNames{
		registry.BEGIN,
		registry.END,
		registry.UID,
		registry.DTSTAMP,
	}
}

func (tC *toDoCalendarComponent) MutuallyExclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{registry.DUE, registry.DURATION_PROP}
}

func (tC *toDoCalendarComponent) MutuallyInclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}

func (tC *toDoCalendarComponent) SerializeToICSFormat(output io.Writer) {
	tC.Begin.ToICalendarPropFormat(output)
	tC.UID.ToICalendarPropFormat(output)
	tC.DateTimeStamp.ToICalendarPropFormat(output)
	for i := 0; i < len(tC.AlarmCalendarComponents); i++ {
		tC.AlarmCalendarComponents[i].SerializeToICSFormat(output)
	}
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}
