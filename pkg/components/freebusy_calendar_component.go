package components

import (
	"io"

	"github.com/vareversat/gics/pkg/types"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
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
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.4
func NewFreeBusyCalendarComponent(
	uid properties.UidProperty,
	dateTimeStamp properties.DateTimeStampProperty,
	propertyList ...properties.Property) FreeBusyCalendarComponent {
	return &freeBusyCalendarComponent{
		Begin:         properties.NewBlockDelimiterProperty(registries.BEGIN, types.VFREEBUSY),
		UID:           uid,
		DateTimeStamp: dateTimeStamp,
		Properties:    propertyList,
		End:           properties.NewBlockDelimiterProperty(registries.END, types.VFREEBUSY),
	}
}

func (fC *freeBusyCalendarComponent) GetProperty(
	name registries.PropertyNames,
) properties.Property {
	for i := 0; i < len(fC.Properties); i++ {
		if fC.Properties[i].GetName() == name {
			return fC.Properties[i]
		}
	}
	return nil
}

func (fC *freeBusyCalendarComponent) MandatoryProperties() []registries.PropertyNames {
	return []registries.PropertyNames{
		registries.BEGIN,
		registries.END,
		registries.UID,
		registries.DTSTAMP,
	}
}

func (fC *freeBusyCalendarComponent) MutuallyExclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{}
}

func (fC *freeBusyCalendarComponent) MutuallyInclusiveProperties() []registries.PropertyNames {
	return []registries.PropertyNames{}
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
