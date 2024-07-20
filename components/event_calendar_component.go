package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
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
			registries.BEGIN,
			registries.Vevent,
		),
		UID:                     uid,
		DateTimeStamp:           dateTimeStamp,
		AlarmCalendarComponents: alarmCalendarComponents,
		Properties:              propertyList,
		End: properties.NewBlockDelimiterProperty(
			registries.END,
			registries.Vevent,
		),
	}
}

func (eC *eventCalendarComponent) GetProperty(name registries.PropertyNames) properties.Property {
	for i := 0; i < len(eC.Properties); i++ {
		if eC.Properties[i].GetName() == name {
			return eC.Properties[i]
		}
	}
	return nil
}
func (eC *eventCalendarComponent) MandatoryProperties() []registries.PropertyNames {
	return []registries.PropertyNames{
		registries.BEGIN,
		registries.END,
		registries.UID,
		registries.DTSTAMP,
	}
}

func (eC *eventCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{registries.DTEND, registries.DURATION_PROP}
}

func (eC *eventCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{}
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
