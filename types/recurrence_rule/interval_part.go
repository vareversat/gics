package recurrence_rule

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type IntervalPart interface {
	RecurrenceRulePart
}

type intervalPart struct {
	partName RecurrenceRulePartName
	interval types.IntegerType
}

// NewIntervalPart give the info of the recurrence occurs repeat interval. See [RFC-5545] ref for more info
// Example: INTERVAL=10 => "every FREQ event based type"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewIntervalPart(interval int32) IntervalPart {
	return &intervalPart{
		partName: Interval,
		interval: types.NewIntegerValue(interval),
	}
}

func (p *intervalPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *intervalPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *intervalPart) GetPartValue() string {
	return p.interval.GetStringValue()
}
