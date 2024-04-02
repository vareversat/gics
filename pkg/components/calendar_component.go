package components

import (
	"io"
)

type CalendarComponent interface {
	ToICalendarComponentFormat(output io.Writer)
}

type Components []CalendarComponent
