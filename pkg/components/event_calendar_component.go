package components

import (
	"io"

	"github.com/vareversat/gics/pkg/types"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

// EventCalendarComponent is the interface definition of a VEVENT calendar component
// See also : [CalendarComponent]
type EventCalendarComponent interface {
	CalendarComponent
}

type eventCalendarComponent struct {
	Begin         properties.BlockDelimiterProperty
	UID           properties.UidProperty
	DateTimeStamp properties.DateTimeStampProperty
	Properties    properties.Properties
	End           properties.BlockDelimiterProperty
}

// NewEventCalendarComponent create a VEVENT calendar component
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.1
func NewEventCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	propertyList ...properties.Property) EventCalendarComponent {
	return &eventCalendarComponent{
		Begin:         properties.NewBlockDelimiterProperty(registries.BEGIN, types.VEVENT),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registries.END, types.VEVENT),
	}
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
	for i := 0; i < len(eC.Properties); i++ {
		eC.Properties[i].ToICalendarPropFormat(output)
	}
	eC.End.ToICalendarPropFormat(output)
}
