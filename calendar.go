package gics

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/types"

	"github.com/vareversat/gics/components"
	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registry"
)

// Calendar is the interface definition of a iCalendar object
type Calendar interface {
	// SerializeToICSFormat format the CalendarComponent to his RFC-5545 representation.
	// It can be print inside a file or the stdout via the "output" param
	SerializeToICSFormat(output io.Writer)
}

type calendar struct {
	Begin      properties.BlockDelimiterProperty // Mandatory
	ProdId     properties.ProductIdProperty      // Mandatory
	Version    properties.VersionProperty        // Mandatory
	Method     properties.MethodProperty         // Optional
	CalScale   properties.CalendarScaleProperty  // Optional
	Components components.CalendarComponents     // At list one
	End        properties.BlockDelimiterProperty // Mandatory
}

// NewCalendar create a iCalendar object
// You can find more information in this section of the RFC-5545 : https://datatracker.ietf.org/doc/html/rfc5545#section-3.4
func NewCalendar(
	calendarComponents components.CalendarComponents,
	prodId string,
	calendarMethod string,
	calendarVersion string,
) (Calendar, error) {
	if calendarComponents == nil || len(calendarComponents) == 0 {
		return nil, fmt.Errorf("you must at least specify a calendar component")
	}
	if calendarVersion == "" {
		return nil, fmt.Errorf("you must specify a VERSION number")
	}
	if prodId == "" {
		return nil, fmt.Errorf("you must specify a PRODID number")
	}
	return &calendar{
		Begin:      properties.NewBlockDelimiterProperty(registry.BEGIN, types.VCALENDAR),
		ProdId:     properties.NewProductIdProperty(prodId),
		Version:    properties.NewVersionProperty(calendarVersion),
		Method:     properties.NewMethodProperty(calendarMethod),
		CalScale:   properties.NewCalScaleProperty(),
		Components: calendarComponents,
		End:        properties.NewBlockDelimiterProperty(registry.END, types.VCALENDAR),
	}, nil
}

func (c *calendar) SerializeToICSFormat(output io.Writer) {
	c.Begin.ToICalendarPropFormat(output)
	c.ProdId.ToICalendarPropFormat(output)
	c.Version.ToICalendarPropFormat(output)
	if c.Method != nil {
		c.Method.ToICalendarPropFormat(output)
	}
	if c.CalScale != nil {
		c.CalScale.ToICalendarPropFormat(output)
	}
	for i := 0; i < len(c.Components); i++ {
		c.Components[i].SerializeToICSFormat(output)
	}
	c.End.ToICalendarPropFormat(output)
}
