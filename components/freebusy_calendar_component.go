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
	Begin      properties.BeginProperty
	Properties properties.Properties
	End        properties.EndProperty
}

// NewFreeBusyCalendarComponent create a VFREEBUSY calendar component
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.4
func NewFreeBusyCalendarComponent(
	propertyList ...properties.Property) FreeBusyCalendarComponent {
	return &freeBusyCalendarComponent{
		Begin: properties.NewBeginProperty(
			registries.Vfreebusy,
		),
		Properties: propertyList,
		End: properties.NewEndProperty(
			registries.Vfreebusy,
		),
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
		registries.BeginProp,
		registries.EndProp,
		registries.UidProp,
		registries.DateTimeStampProp,
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
	for i := 0; i < len(fC.Properties); i++ {
		fC.Properties[i].ToICalendarPropFormat(output)
	}
	fC.End.ToICalendarPropFormat(output)
}
