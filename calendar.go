package gics

import (
	"io"

	"github.com/vareversat/gics/components"
	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
)

// Calendar is the interface definition of a iCalendar object
type Calendar interface {
	// SerializeToICSFormat format the CalendarComponent to his RFC-5545 representation.
	// It can be print inside a file or the stdout via the "output" param
	SerializeToICSFormat(output io.Writer)

	// MandatoryProperties return the list of the mandatory properties of the Calendar
	MandatoryProperties() []registries.PropertyRegistry

	// GetProperty get a property by his registries.PropertyRegistry
	GetProperty(name registries.PropertyRegistry) properties.Property

	// AddProperty add a new property to a Calendar
	AddProperty(property properties.Property)
}

type calendar struct {
	Begin      properties.BeginProperty
	Properties properties.Properties
	Component  components.Component
	End        properties.EndProperty
}

// NewCalendar create a iCalendar object
// You can find more information in this section of the RFC-5545 : https://datatracker.ietf.org/doc/html/rfc5545#section-3.4
func NewCalendar(
	calendarComponent components.Component,
	propertyList ...properties.Property,
) (Calendar, error) {
	return &calendar{
		Begin: properties.NewBeginProperty(
			registries.Vcalendar,
		),
		Properties: propertyList,
		Component:  calendarComponent,
		End: properties.NewEndProperty(
			registries.Vcalendar,
		),
	}, nil
}

func (c *calendar) GetProperty(
	name registries.PropertyRegistry,
) properties.Property {
	for i := 0; i < len(c.Properties); i++ {
		if c.Properties[i].GetName() == name {
			return c.Properties[i]
		}
	}
	return nil
}
func (c *calendar) MandatoryProperties() []registries.PropertyRegistry {
	return []registries.PropertyRegistry{
		registries.BeginProp,
		registries.EndProp,
		registries.ProductIdentifierProp,
		registries.VersionProp,
	}
}

func (c *calendar) AddProperty(property properties.Property) {
	c.Properties = append(c.Properties, property)
}

func (c *calendar) SerializeToICSFormat(output io.Writer) {
	c.Begin.ToICalendarPropFormat(output)
	for i := 0; i < len(c.Properties); i++ {
		c.Properties[i].ToICalendarPropFormat(output)
	}
	c.Component.SerializeToICSFormat(output)
	c.End.ToICalendarPropFormat(output)
}
