package components

import (
	"io"

	"github.com/vareversat/gics/types"

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
		Begin:      properties.NewBlockDelimiterProperty(registries.BEGIN, types.VALARM),
		Trigger:    trigger,
		Properties: propertyList,
		Action:     action,
		End:        properties.NewBlockDelimiterProperty(registries.END, types.VALARM),
	}
}

func (aC *alarmCalendarComponent) GetProperty(name registries.PropertyNames) properties.Property {
	for i := 0; i < len(aC.Properties); i++ {
		if aC.Properties[i].GetName() == name {
			return aC.Properties[i]
		}
	}
	return nil
}

func (aC *alarmCalendarComponent) MandatoryProperties() []registries.PropertyNames {
	switch aC.Action.GetActionValue() {
	case types.AUDIO:
		return []registries.PropertyNames{
			registries.BEGIN,
			registries.END,
			registries.ACTION,
			registries.TRIGGER,
		}
	case types.DISPLAY:
		return []registries.PropertyNames{
			registries.BEGIN,
			registries.END,
			registries.ACTION,
			registries.TRIGGER,
			registries.DESCRIPTION,
		}
	case types.EMAIL:
		return []registries.PropertyNames{
			registries.BEGIN,
			registries.END,
			registries.ACTION,
			registries.TRIGGER,
			registries.DESCRIPTION,
			registries.SUMMARY,
		}
	default:
		return []registries.PropertyNames{registries.BEGIN, registries.END}
	}
}

func (aC *alarmCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{}
}

func (aC *alarmCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{registries.DURATION_PROP, registries.REPEAT}
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
