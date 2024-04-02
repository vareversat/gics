package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.5

import (
	"io"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type FreeBusyCalendarComponent interface {
	CalendarComponent
}

type freeBusyCalendarComponent struct {
	Begin properties.BlockDelimiterProperty
	End   properties.BlockDelimiterProperty

	Properties properties.Properties
}

func NewFreeBusyCalendarComponent(propertyList ...properties.Property) FreeBusyCalendarComponent {
	return &freeBusyCalendarComponent{
		Begin:      properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VFREEBUSY),
		End:        properties.NewBlockDelimiterProperty(registries.END, properties.VFREEBUSY),
		Properties: propertyList,
	}
}

func (fC *freeBusyCalendarComponent) ToICalendarComponentFormat(output io.Writer) {
	fC.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(fC.Properties); i++ {
		fC.Properties[i].ToICalendarPropFormat(output)
	}
	fC.End.ToICalendarPropFormat(output)
}
