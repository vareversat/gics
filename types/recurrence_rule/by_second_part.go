package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

const (
	minSecond = 0
	maxSecond = 60
)

type BySecondPart interface {
	RecurrenceRulePart
}

type bySecondPart struct {
	partName RecurrenceRulePartName
	seconds  []types.IntegerType
}

// NewBySecondPart give the info on which second of the minute the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYSECOND=2 => "2nd second of the minute"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewBySecondPart(seconds []int32) (BySecondPart, error) {
	for i := 0; i < len(seconds); i++ {
		if seconds[i] < minMonth || seconds[i] > maxMonth {
			return &bySecondPart{
				partName: ByMonth,
				seconds:  types.NewIntegerValues([]int32{0}),
			}, fmt.Errorf("%d (%d-th) is not a valid second. It must satisfy this : second ∈ [%d;%d]", seconds[i], i, minSecond, maxSecond)
		}
	}
	return &bySecondPart{
		partName: ByMonth,
		seconds:  types.NewIntegerValues(seconds),
	}, nil
}

func (p *bySecondPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *bySecondPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *bySecondPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.seconds); i++ {
		secondsOutput.Write([]byte(p.seconds[i].GetStringValue()))
		if len(p.seconds)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
