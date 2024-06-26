package components

import (
	"io"

	"github.com/vareversat/gics/pkg/registries"
)

// CalendarComponent is the common interface definition share among VEVENT, VALARM, VJOURNAL, VTODO, VFREEBUSY & VTIMEZONE
// You can find more information in this section of the RFC-5545 : https://datatracker.ietf.org/doc/html/rfc5545#section-3.6
type CalendarComponent interface {
	// SerializeToICSFormat format the CalendarComponent to his RFC-5545 representation.
	// It can be print inside a file or the stdout via the "output" param
	SerializeToICSFormat(output io.Writer)

	// MandatoryProperties return the list of the mandatory properties of a CalendarComponent
	MandatoryProperties() []registries.PropertyNames

	// MutuallyExclusiveProperties return the list of the mutually exclusives properties of a CalendarComponent
	// Example : In a VEVENT component, you can have a DTEND and a DURATION property at the same time
	MutuallyExclusiveProperties() []registries.PropertyNames

	// MutuallyInclusiveProperties return the list of the mutually inclusive properties of a CalendarComponent
	// Example : In a VALARM component, if you set a value for the DURATION property, you have to also set one for the REPEAT property
	MutuallyInclusiveProperties() []registries.PropertyNames
}

// CalendarComponents is an array of CalendarComponent
type CalendarComponents []CalendarComponent
