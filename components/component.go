package components

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/properties"

	"github.com/vareversat/gics/registries"
)

// Component is the common interface definition share among VEVENT, VALARM, VJOURNAL, VTODO, VFREEBUSY & VTIMEZONE
// You can find more information in this section of the RFC-5545 : https://datatracker.ietf.org/doc/html/rfc5545#section-3.6
type Component interface {
	// SerializeToICSFormat format the Component to his RFC-5545 representation.
	// It can be print inside a file or the stdout via the "output" param
	SerializeToICSFormat(output io.Writer)

	// MandatoryProperties return the list of the mandatory properties of a Component
	MandatoryProperties() []registries.PropertyRegistry

	// GetProperty get a property by his registries.PropertyRegistry
	GetProperty(name registries.PropertyRegistry) properties.Property

	// AddProperty add a new property to a component
	AddProperty(property properties.Property)

	// MutuallyExclusiveProperties return the list of the mutually exclusives properties of a Component
	// Example : In a VEVENT component, you can't have a DTEND and a DURATION property at the same time
	MutuallyExclusiveProperties() []registries.PropertyRegistry

	// MutuallyInclusiveProperties return the list of the mutually inclusive properties of a Component
	// Example : In a VALARM component, if you set a value for the DURATION property, you have to also set one for the REPEAT property
	MutuallyInclusiveProperties() []registries.PropertyRegistry
}

// Components is an array of Component
type Components []Component

// NewComponentFromName create a new Component based on the first BEGIN: block
func NewComponentFromName(name registries.ComponentRegistry) (Component, error) {
	switch name {
	case registries.Valarm:
		return NewAlarmCalendarComponent(nil), nil
	case registries.Vevent:
		return NewEventCalendarComponent(nil), nil
	case registries.Vfreebusy:
		return NewFreeBusyCalendarComponent(nil), nil
	case registries.Vjournal:
		return NewJournalCalendarComponent(nil), nil
	case registries.Vtimezone:
		return NewTimeZoneCalendarComponent(nil), nil
	case registries.Vtodo:
		return NewToDoCalendarComponent(nil), nil
	default:
		return nil, fmt.Errorf("%s is not a valid calendar component name", string(name))
	}
}
