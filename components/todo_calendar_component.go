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
	// AddAlarmComponent allow to add an AlarmCalendarComponent to the ToDoCalendarComponent
	AddAlarmComponent(alarm AlarmCalendarComponent)
}

type toDoCalendarComponent struct {
	Begin                   properties.BeginProperty
	AlarmCalendarComponents []AlarmCalendarComponent
	Properties              properties.Properties
	End                     properties.EndProperty
}

// NewToDoCalendarComponent create a VTODO calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.2
func NewToDoCalendarComponent(
	alarmCalendarComponents []AlarmCalendarComponent,
	propertyList ...properties.Property) ToDoCalendarComponent {
	return &toDoCalendarComponent{
		Begin: properties.NewBeginProperty(
			registries.Vtodo,
		),
		AlarmCalendarComponents: alarmCalendarComponents,
		Properties:              propertyList,
		End: properties.NewEndProperty(
			registries.Vtodo,
		),
	}
}

func (tC *toDoCalendarComponent) AddAlarmComponent(alarm AlarmCalendarComponent) {
	tC.AlarmCalendarComponents = append(tC.AlarmCalendarComponents, alarm)
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
	for i := 0; i < len(tC.Properties); i++ {
		tC.Properties[i].ToICalendarPropFormat(output)
	}
	for i := 0; i < len(tC.AlarmCalendarComponents); i++ {
		tC.AlarmCalendarComponents[i].SerializeToICSFormat(output)
	}
	tC.End.ToICalendarPropFormat(output)
}
