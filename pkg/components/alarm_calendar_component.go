package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6

import (
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

func (aC *alarmCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	aC.Begin.ToICalendarPropFormat(output)
	aC.Action.ToICalendarPropFormat(output)
	aC.Trigger.ToICalendarPropFormat(output)
	for i := 0; i < len(aC.Properties); i++ {
		aC.Properties[i].ToICalendarPropFormat(output)
	}
	aC.End.ToICalendarPropFormat(output)
}
