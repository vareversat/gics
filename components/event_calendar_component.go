package components

import (
	"io"

	"github.com/vareversat/gics/types"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registry"
)

// EventCalendarComponent is the interface definition of a VEVENT calendar component
// See also : [CalendarComponent]
type EventCalendarComponent interface {
	CalendarComponent
}

type eventCalendarComponent struct {
	Begin                   properties.BlockDelimiterProperty
	UID                     properties.UidProperty
	DateTimeStamp           properties.DateTimeStampProperty
	AlarmCalendarComponents []AlarmCalendarComponent
	Properties              properties.Properties
	End                     properties.BlockDelimiterProperty
}

// NewEventCalendarComponent create a VEVENT calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.1
func NewEventCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	alarmCalendarComponents []AlarmCalendarComponent,
	propertyList ...properties.Property) EventCalendarComponent {
	return &eventCalendarComponent{
		Begin: properties.NewBlockDelimiterProperty(
			registry.BEGIN,
			types.VEVENT,
		),
		UID:                     uid,
		DateTimeStamp:           dateTimeStamp,
		AlarmCalendarComponents: alarmCalendarComponents,
		Properties:              propertyList,
		End:                     properties.NewBlockDelimiterProperty(registry.END, types.VEVENT),
	}
}

func (eC *eventCalendarComponent) GetProperty(name registry.PropertyNames) properties.Property {
	for i := 0; i < len(eC.Properties); i++ {
		if eC.Properties[i].GetName() == name {
			return eC.Properties[i]
		}
	}
	return nil
}
func (eC *eventCalendarComponent) MandatoryProperties() []registry.PropertyNames {
	return []registry.PropertyNames{
		registry.BEGIN,
		registry.END,
		registry.UID,
		registry.DTSTAMP,
	}
}

func (eC *eventCalendarComponent) MutuallyExclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{registry.DTEND, registry.DURATION_PROP}
}

func (eC *eventCalendarComponent) MutuallyInclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}

func (eC *eventCalendarComponent) SerializeToICSFormat(output io.Writer) {
	eC.Begin.ToICalendarPropFormat(output)
	eC.UID.ToICalendarPropFormat(output)
	eC.DateTimeStamp.ToICalendarPropFormat(output)
	for i := 0; i < len(eC.AlarmCalendarComponents); i++ {
		eC.AlarmCalendarComponents[i].SerializeToICSFormat(output)
	}
	for i := 0; i < len(eC.Properties); i++ {
		eC.Properties[i].ToICalendarPropFormat(output)
	}
	eC.End.ToICalendarPropFormat(output)
}
