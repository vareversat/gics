package components

import (
	"io"

	"github.com/vareversat/gics/properties"

	"github.com/vareversat/gics/registries"
)

// CalendarComponent is the common interface definition share among VEVENT, VALARM, VJOURNAL, VTODO, VFREEBUSY & VTIMEZONE
// You can find more information in this section of the RFC-5545 : https://datatracker.ietf.org/doc/html/rfc5545#section-3.6
type CalendarComponent interface {
	// SerializeToICSFormat format the CalendarComponent to his RFC-5545 representation.
	// It can be print inside a file or the stdout via the "output" param
	SerializeToICSFormat(output io.Writer)

	// MandatoryProperties return the list of the mandatory properties of a CalendarComponent
	MandatoryProperties() []registries.PropertyRegistry

	// GetProperty get a property by his registries.PropertyRegistry
	GetProperty(name registries.PropertyRegistry) properties.Property

	// MutuallyExclusiveProperties return the list of the mutually exclusives properties of a CalendarComponent
	// Example : In a VEVENT component, you can't have a DTEND and a DURATION property at the same time
	MutuallyExclusiveProperties() []registries.PropertyRegistry

	// MutuallyInclusiveProperties return the list of the mutually inclusive properties of a CalendarComponent
	// Example : In a VALARM component, if you set a value for the DURATION property, you have to also set one for the REPEAT property
	MutuallyInclusiveProperties() []registries.PropertyRegistry
}

// CalendarComponents is an array of CalendarComponent
type CalendarComponents []CalendarComponent
