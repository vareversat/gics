package recurrence_rule

import (
	"fmt"
	"io"
	"time"

	"github.com/vareversat/gics/types"
)

type UntilPart interface {
	RecurrenceRulePart
}

type untilPart struct {
	partName RecurrenceRulePartName
	endDate  types.DateTimeType
}

// NewUntilPart give info of the bounds the recurrence rule in an inclusive manner. See [RFC-5545] ref for more info
// Example: UNTIL=19730429T070000Z => "until April 29, 1973 at 07:00:00 UTC"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewUntilPart(endDate time.Time) UntilPart {
	return &untilPart{
		partName: Until,
		endDate:  types.NewDateTimeValue(endDate),
	}
}

func (p *untilPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *untilPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *untilPart) GetPartValue() string {
	return p.endDate.GetStringValue()
}
