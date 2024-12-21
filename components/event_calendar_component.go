package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
)

// EventCalendarComponent is the interface definition of a VEVENT calendar component
// See also : [Component]
type EventCalendarComponent interface {
	Component
	// AddAlarmComponent allow to add an AlarmCalendarComponent to the EventCalendarComponent
	AddAlarmComponent(alarm AlarmCalendarComponent)
}

type eventCalendarComponent struct {
	Begin                   properties.BeginProperty
	AlarmCalendarComponents []AlarmCalendarComponent
	Properties              properties.Properties
	End                     properties.EndProperty
}

// NewEventCalendarComponent create a VEVENT calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.1
func NewEventCalendarComponent(
	propertyList ...properties.Property) EventCalendarComponent {
	return &eventCalendarComponent{
		Begin: properties.NewBeginProperty(
			registries.Vevent,
		),
		AlarmCalendarComponents: []AlarmCalendarComponent{},
		Properties:              propertyList,
		End: properties.NewEndProperty(
			registries.Vevent,
		),
	}
}

func (eC *eventCalendarComponent) AddAlarmComponent(alarm AlarmCalendarComponent) {
	eC.AlarmCalendarComponents = append(eC.AlarmCalendarComponents, alarm)
}

func (eC *eventCalendarComponent) GetProperty(
	name registries.PropertyRegistry,
) properties.Property {
	for i := 0; i < len(eC.Properties); i++ {
		if eC.Properties[i].GetName() == name {
			return eC.Properties[i]
		}
	}
	return nil
}
func (eC *eventCalendarComponent) MandatoryProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{
		registries.BeginProp,
		registries.EndProp,
		registries.UidProp,
		registries.DateTimeStampProp,
	}
}

func (eC *eventCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{registries.DateTimeEndProp, registries.DurationProp}
}

func (eC *eventCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (eC *eventCalendarComponent) AddProperty(property properties.Property) {
	eC.Properties = append(eC.Properties, property)
}

func (eC *eventCalendarComponent) SerializeToICSFormat(output io.Writer) {
	eC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(eC.Properties); i++ {
		eC.Properties[i].ToICalendarPropFormat(output)
	}
	for i := 0; i < len(eC.AlarmCalendarComponents); i++ {
		eC.AlarmCalendarComponents[i].SerializeToICSFormat(output)
	}
	eC.End.ToICalendarPropFormat(output)
}
