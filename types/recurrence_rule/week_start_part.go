package recurrence_rule

import (
	"fmt"
	"io"
)

type WeekStartPart interface {
	RecurrenceRulePart
}

type weekStartPart struct {
	partName RecurrenceRulePartName
	weekDay  WeekDay
}

// NewWeekStartPart give info of the week start. See [RFC-5545] ref for more info
// Example: WKST=MO => "week start on monday"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewWeekStartPart(weekday WeekDay) WeekStartPart {
	return &weekStartPart{
		partName: WeekStart,
		weekDay:  weekday,
	}
}

func (p *weekStartPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *weekStartPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *weekStartPart) GetPartValue() string {
	return string(p.weekDay)
}
