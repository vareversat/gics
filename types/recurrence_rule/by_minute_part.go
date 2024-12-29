package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

const (
	minMinute = 0
	maxMinute = 59
)

type ByMinutePart interface {
	RecurrenceRulePart
}

type byMinutePart struct {
	partName RecurrenceRulePartName
	minutes  []types.IntegerType
}

// NewByMinutePart give the info on which hour of the day the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYMINUTE=19 => "at XX:19 (XX:19 am/pm)"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByMinutePart(minutes []int32) (ByMinutePart, error) {
	for i := 0; i < len(minutes); i++ {
		if minutes[i] < minMinute || minutes[i] > maxMinute {
			return &byMinutePart{
				partName: ByMinute,
				minutes:  types.NewIntegerValues([]int32{0}),
			}, fmt.Errorf("%d (%d-th) is not a valid minute. It must satisfy this : minute ∈ [%d;%d]", minutes[i], i, minMinute, maxMinute)
		}
	}
	return &byMinutePart{
		partName: ByMinute,
		minutes:  types.NewIntegerValues(minutes),
	}, nil
}

func (p *byMinutePart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *byMinutePart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *byMinutePart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.minutes); i++ {
		secondsOutput.Write([]byte(p.minutes[i].GetStringValue()))
		if len(p.minutes)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
