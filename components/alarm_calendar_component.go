package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
)

// AlarmCalendarComponent is the interface definition of a VALARM calendar component
// See also : [Component]
type AlarmCalendarComponent interface {
	Component
}

type alarmCalendarComponent struct {
	Begin      properties.BeginProperty
	Properties properties.Properties
	End        properties.EndProperty
}

// NewAlarmCalendarComponent create a VALARM calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6
func NewAlarmCalendarComponent(
	propertyList ...properties.Property) AlarmCalendarComponent {
	return &alarmCalendarComponent{
		Begin:      properties.NewBeginProperty(registries.Valarm),
		Properties: propertyList,
		End:        properties.NewEndProperty(registries.Valarm),
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
	actionProp := aC.GetProperty(registries.ActionProp)
	if actionProp != nil {
		switch registries.ActionRegistry(actionProp.GetValue()) {
		case registries.Audio:
			return []registries.PropertyRegistry{
				registries.BeginProp,
				registries.EndProp,
				registries.ActionProp,
				registries.TriggerProp,
			}
		case registries.Display:
			return []registries.PropertyRegistry{
				registries.BeginProp,
				registries.EndProp,
				registries.ActionProp,
				registries.TriggerProp,
				registries.DescriptionProp,
			}
		case registries.Email:
			return []registries.PropertyRegistry{
				registries.BeginProp,
				registries.EndProp,
				registries.ActionProp,
				registries.TriggerProp,
				registries.DescriptionProp,
				registries.SummaryProp,
			}
		default:
			return []registries.PropertyRegistry{registries.BeginProp, registries.EndProp}

		}
	} else {
		return []registries.PropertyRegistry{}
	}
}

func (aC *alarmCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (aC *alarmCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{registries.DurationProp, registries.RepeatProp}
}

func (aC *alarmCalendarComponent) AddProperty(property properties.Property) {
	aC.Properties = append(aC.Properties, property)
}

func (aC *alarmCalendarComponent) SerializeToICSFormat(output io.Writer) {
	aC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(aC.Properties); i++ {
		aC.Properties[i].ToICalendarPropFormat(output)
	}
	aC.End.ToICalendarPropFormat(output)
}
