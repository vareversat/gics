package components

import (
	"io"

	"github.com/vareversat/gics/types"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registry"
)

// FreeBusyCalendarComponent is the interface definition of a VFREEBUSY calendar component
// See also : [CalendarComponent]
type FreeBusyCalendarComponent interface {
	CalendarComponent
}

type freeBusyCalendarComponent struct {
	Begin         properties.BlockDelimiterProperty
	UID           properties.UidProperty
	DateTimeStamp properties.DateTimeStampProperty
	Properties    properties.Properties
	End           properties.BlockDelimiterProperty
}

// NewFreeBusyCalendarComponent create a VFREEBUSY calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.4
func NewFreeBusyCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	propertyList ...properties.Property) FreeBusyCalendarComponent {
	return &freeBusyCalendarComponent{
		Begin:         properties.NewBlockDelimiterProperty(registry.BEGIN, types.VFREEBUSY),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registry.END, types.VFREEBUSY),
	}
}

func (fC *freeBusyCalendarComponent) GetProperty(
	name registry.PropertyNames,
) properties.Property {
	for i := 0; i < len(fC.Properties); i++ {
		if fC.Properties[i].GetName() == name {
			return fC.Properties[i]
		}
	}
	return nil
}

func (fC *freeBusyCalendarComponent) MandatoryProperties() []registry.PropertyNames {
	return []registry.PropertyNames{
		registry.BEGIN,
		registry.END,
		registry.UID,
		registry.DTSTAMP,
	}
}

func (fC *freeBusyCalendarComponent) MutuallyExclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}

func (fC *freeBusyCalendarComponent) MutuallyInclusiveProperties() []registry.PropertyNames {
	return []registry.PropertyNames{}
}

func (fC *freeBusyCalendarComponent) SerializeToICSFormat(output io.Writer) {
	fC.Begin.ToICalendarPropFormat(output)
	fC.UID.ToICalendarPropFormat(output)
	fC.DateTimeStamp.ToICalendarPropFormat(output)
	for i := 0; i < len(fC.Properties); i++ {
		fC.Properties[i].ToICalendarPropFormat(output)
	}
	fC.End.ToICalendarPropFormat(output)
}
