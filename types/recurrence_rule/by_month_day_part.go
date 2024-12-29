package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

const (
	minMonthDay = -31
	maxMonthDay = 31
)

type ByMonthDayPart interface {
	RecurrenceRulePart
}

type byMonthDayPart struct {
	partName RecurrenceRulePartName
	monthDay []types.IntegerType
}

// NewByMonthDayPart give the info on which day of the month the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYMONTHDAY=19 => "19th day of the month"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByMonthDayPart(months []int32) (ByMonthDayPart, error) {
	for i := 0; i < len(months); i++ {
		if months[i] < minMonthDay || months[i] > maxMonthDay || months[i] == 0 {
			return &byMonthDayPart{
				partName: ByMonthDay,
				monthDay: types.NewIntegerValues([]int32{0}),
			}, fmt.Errorf("%d (%d-th) is not a valid day num. It must satisfy this : day num ∈ [%d;0[U]0;%d]", months[i], i, minMonthDay, maxMonthDay)
		}
	}
	return &byMonthDayPart{
		partName: ByMonthDay,
		monthDay: types.NewIntegerValues(months),
	}, nil
}

func (p *byMonthDayPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *byMonthDayPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *byMonthDayPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.monthDay); i++ {
		secondsOutput.Write([]byte(p.monthDay[i].GetStringValue()))
		if len(p.monthDay)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
