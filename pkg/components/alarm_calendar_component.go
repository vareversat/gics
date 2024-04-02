package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6

import (
	"github.com/vareversat/gics/pkg/types"
	"io"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type AlarmCalendarComponent interface {
	CalendarComponent
}

type AlarmCalendarSubComponent interface {
	CalendarComponent
}

type alarmCalendarComponent struct {
	Begin      properties.BlockDelimiterProperty
	Action     properties.ActionProperty
	Trigger    properties.TriggerProperty
	Properties properties.Properties
	End        properties.BlockDelimiterProperty
}

func NewAlarmCalendarComponent(
	action properties.ActionProperty,
	trigger properties.TriggerProperty,
	propertyList ...properties.Property) AlarmCalendarComponent {
	return &alarmCalendarComponent{
		Begin:      properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VALARM),
		Trigger:    trigger,
		Properties: propertyList,
		Action:     action,
		End:        properties.NewBlockDelimiterProperty(registries.END, properties.VALARM),
	}
}
func (aC *alarmCalendarComponent) MandatoryProperties() []registries.PropertyNames {
	switch aC.Action.GetValue() {
	case types.AUDIO:
		return []registries.PropertyNames{registries.BEGIN, registries.END, registries.ACTION, registries.TRIGGER}
	case types.DISPLAY:
		return []registries.PropertyNames{registries.BEGIN, registries.END, registries.ACTION, registries.TRIGGER, registries.DESCRIPTION}
	case types.EMAIL:
		return []registries.PropertyNames{registries.BEGIN, registries.END, registries.ACTION, registries.TRIGGER, registries.DESCRIPTION, registries.SUMMARY}
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

func (aC *alarmCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	aC.Begin.ToICalendarPropFormat(output)
	aC.Action.ToICalendarPropFormat(output)
	aC.Trigger.ToICalendarPropFormat(output)
	for i := 0; i < len(aC.Properties); i++ {
		aC.Properties[i].ToICalendarPropFormat(output)
	}
	aC.End.ToICalendarPropFormat(output)
}
