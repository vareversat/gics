package components

import (
	"github.com/vareversat/gics/pkg/registries"
	"io"
)

type CalendarComponent interface {
	ToICalendarComponentFormat(output io.Writer)
	MandatoryProperties() []registries.PropertyNames
	MutuallyExclusiveProperties() []registries.PropertyNames
	MutuallyInclusiveProperties() []registries.PropertyNames
}

type Components []CalendarComponent
