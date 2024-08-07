package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
)

// AlarmCalendarComponent is the interface definition of a VALARM calendar component
// See also : [CalendarComponent]
type AlarmCalendarComponent interface {
	CalendarComponent
}

type alarmCalendarComponent struct {
	Begin      properties.BlockDelimiterProperty
	Action     properties.ActionProperty
	Trigger    properties.TriggerProperty
	Properties properties.Properties
	End        properties.BlockDelimiterProperty
}

// NewAlarmCalendarComponent create a VALARM calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6
func NewAlarmCalendarComponent(
	action properties.ActionProperty,
	trigger properties.TriggerProperty,
	propertyList ...properties.Property) AlarmCalendarComponent {
	return &alarmCalendarComponent{
		Begin:      properties.NewBlockDelimiterProperty(registries.Begin, registries.Valarm),
		Trigger:    trigger,
		Properties: propertyList,
		Action:     action,
		End:        properties.NewBlockDelimiterProperty(registries.End, registries.Valarm),
	}
}

func (aC *alarmCalendarComponent) GetProperty(
	name registries.PropertyRegistry,
) properties.Property {
	for i := 0; i < len(aC.Properties); i++ {
		if aC.Properties[i].GetName() == name {
			return aC.Properties[i]
		}
	}
	return nil
}

func (aC *alarmCalendarComponent) MandatoryProperties() []registries.PropertyRegistry {
	switch aC.Action.GetActionValue() {
	case registries.Audio:
		return []registries.PropertyRegistry{
			registries.Begin,
			registries.End,
			registries.Action,
			registries.Trigger,
		}
	case registries.Display:
		return []registries.PropertyRegistry{
			registries.Begin,
			registries.End,
			registries.Action,
			registries.Trigger,
			registries.Description,
		}
	case registries.Email:
		return []registries.PropertyRegistry{
			registries.Begin,
			registries.End,
			registries.Action,
			registries.Trigger,
			registries.Description,
			registries.Summary,
		}
	default:
		return []registries.PropertyRegistry{registries.Begin, registries.End}
	}
}

func (aC *alarmCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (aC *alarmCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{registries.DurationProperty, registries.Repeat}
}

func (aC *alarmCalendarComponent) SerializeToICSFormat(output io.Writer) {
	aC.Begin.ToICalendarPropFormat(output)
	aC.Action.ToICalendarPropFormat(output)
	aC.Trigger.ToICalendarPropFormat(output)
	for i := 0; i < len(aC.Properties); i++ {
		aC.Properties[i].ToICalendarPropFormat(output)
	}
	aC.End.ToICalendarPropFormat(output)
}
