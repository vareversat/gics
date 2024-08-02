package components

import (
	"io"

	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
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
		Begin:         properties.NewBlockDelimiterProperty(registries.Begin, registries.Vfreebusy),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registries.End, registries.Vfreebusy),
	}
}

func (fC *freeBusyCalendarComponent) GetProperty(
	name registries.PropertyRegistry,
) properties.Property {
	for i := 0; i < len(fC.Properties); i++ {
		if fC.Properties[i].GetName() == name {
			return fC.Properties[i]
		}
	}
	return nil
}

func (fC *freeBusyCalendarComponent) MandatoryProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{
		registries.Begin,
		registries.End,
		registries.Uid,
		registries.DateTimeStamp,
	}
}

func (fC *freeBusyCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
}

func (fC *freeBusyCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{}
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
