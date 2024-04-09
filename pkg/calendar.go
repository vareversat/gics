package pkg

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.4

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/pkg/types"

	"github.com/vareversat/gics/pkg/components"
	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type Calendar interface {
	ToICalendarFormat(output io.Writer)
}

type calendar struct {
	Begin      properties.BlockDelimiterProperty // Mandatory
	ProdId     properties.ProductIdProperty      // Mandatory
	Version    properties.VersionProperty        // Mandatory
	Method     properties.MethodProperty         // Optional
	CalScale   properties.CalendarScaleProperty  // Optional
	Components components.Components             // At list one
	End        properties.BlockDelimiterProperty // Mandatory
}

func NewCalendar(
	calendarComponents components.Components,
	proId string,
	calendarMethod string,
	calendarVersion string,
) (Calendar, error) {
	if calendarComponents == nil || len(calendarComponents) == 0 {
		return nil, fmt.Errorf("you must at least specify a calendar component")
	}
	if calendarVersion == "" {
		return nil, fmt.Errorf("you must specify a VERSION number")
	}
	if proId == "" {
		return nil, fmt.Errorf("you must specify a PRODID number")
	}
	return &calendar{
		Begin:      properties.NewBlockDelimiterProperty(registries.BEGIN, types.VCALENDAR),
		ProdId:     properties.NewProductIdProperty(proId),
		Version:    properties.NewVersionProperty(calendarVersion),
		Method:     properties.NewMethodProperty(calendarMethod),
		CalScale:   properties.NewCalScaleProperty(),
		Components: calendarComponents,
		End:        properties.NewBlockDelimiterProperty(registries.END, types.VCALENDAR),
	}, nil
}

func (c *calendar) ToICalendarFormat(output io.Writer) {
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
		c.Components[i].ToICalendarComponentFormat(output)
	}
	c.End.ToICalendarPropFormat(output)
}
