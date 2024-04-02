package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6

import (
	"io"

	"github.com/vareversat/gics/pkg/types"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type AlarmCalendarComponent interface {
	CalendarComponent
}

type alarmCalendarComponent struct {
	// MANDATORY and MUST NOT occur more than once
	Begin   properties.BlockDelimiterProperty
	End     properties.BlockDelimiterProperty
	Action  properties.ActionProperty
	Trigger properties.TriggerProperty
}

type audioBlock struct {
	alarmCalendarComponent
	// EITHER both of them or NEITHER
	Duration    properties.DurationProperty
	RepeatCount properties.RepeatCountProperty

	// OPTIONAL and MUST NOT occur more than once
	Attachment properties.AttachmentProperty
}

type displayBlock struct {
	alarmCalendarComponent
	// MANDATORY and MUST NOT occur more than once
	Description properties.DescriptionProperty

	// EITHER both of them or NEITHER
	Duration    properties.DurationProperty
	RepeatCount properties.RepeatCountProperty

	// OPTIONAL and MUST NOT occur more than once
	Attachment properties.AttachmentProperty
}

type emailBlock struct {
	alarmCalendarComponent
	// MANDATORY and MUST NOT occur more than once
	Description properties.DescriptionProperty
	Summary     properties.SummaryProperty

	// MANDATORY and MAY occur more than once
	Attendees []properties.AttendeeProperty

	// EITHER both of them or NEITHER
	Duration    properties.DurationProperty
	RepeatCount properties.RepeatCountProperty

	// OPTIONAL and MUST NOT occur more than once
	Attachment properties.AttachmentProperty
}

func NewAlarmCalendarComponent(action types.ActionType) AlarmCalendarComponent {
	switch action {
	case types.AUDIO:
		return &audioBlock{
			alarmCalendarComponent: alarmCalendarComponent{
				Begin:  properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VALARM),
				Action: properties.NewActionProperty(types.AUDIO),
				End:    properties.NewBlockDelimiterProperty(registries.END, properties.VALARM)},
		}
	case types.DISPLAY:
		return &displayBlock{
			alarmCalendarComponent: alarmCalendarComponent{
				Begin:  properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VALARM),
				Action: properties.NewActionProperty(types.DISPLAY),
				End:    properties.NewBlockDelimiterProperty(registries.END, properties.VALARM)},
		}
	case types.EMAIL:
		return &emailBlock{
			alarmCalendarComponent: alarmCalendarComponent{
				Begin:  properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VALARM),
				Action: properties.NewActionProperty(types.EMAIL),
				End:    properties.NewBlockDelimiterProperty(registries.END, properties.VALARM)},
		}
	default:
		return nil
	}
}

func (aC *alarmCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	aC.Begin.ToICalendarPropFormat(output)
	aC.End.ToICalendarPropFormat(output)
}
