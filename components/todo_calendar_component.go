package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
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
			registries.BeginProp,
			registries.Vtodo,
		),
		UID:                     uid,
		DateTimeStamp:           dateTimeStamp,
		AlarmCalendarComponents: alarmCalendarComponents,
		Properties:              propertyList,
		End: properties.NewBlockDelimiterProperty(
			registries.EndProp,
			registries.Vtodo,
		),
	}
}

func (tC *toDoCalendarComponent) GetProperty(name registries.PropertyRegistry) properties.Property {
	for i := 0; i < len(tC.Properties); i++ {
		if tC.Properties[i].GetName() == name {
			return tC.Properties[i]
		}
	}
	return nil
}

func (tC *toDoCalendarComponent) MandatoryProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{
		registries.BeginProp,
		registries.EndProp,
		registries.UidProp,
		registries.DateTimeStampProp,
	}
}

func (tC *toDoCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{registries.DateTimeDueProp, registries.DurationProp}
}

func (tC *toDoCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
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
