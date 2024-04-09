package components

import (
	"io"

	"github.com/vareversat/gics/pkg/registries"
)

type CalendarComponent interface {
	ToICalendarComponentFormat(output io.Writer)
	MandatoryProperties() []registries.PropertyNames
	MutuallyExclusiveProperties() []registries.PropertyNames
	MutuallyInclusiveProperties() []registries.PropertyNames
}

type Components []CalendarComponent
