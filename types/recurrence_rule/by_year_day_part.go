package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

const (
	minDay = -366
	maxDay = 366
)

type ByYearDayPart interface {
	RecurrenceRulePart
}

type byYearDayPart struct {
	partName RecurrenceRulePartName
	days     []types.IntegerType
}

// NewByYearDayPart give the info on which day of the year the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYYEARDAY=100 => "100th day of the year"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByYearDayPart(days []int32) (ByYearDayPart, error) {
	for i := 0; i < len(days); i++ {
		if days[i] < minDay || days[i] > maxDay || days[i] == 0 {
			return &byYearDayPart{
				partName: ByYearDay,
				days:     types.NewIntegerValues([]int32{0}),
			}, fmt.Errorf("%d (%d-th) is not a valid set position. It must satisfy this : day ∈ [%d;0[U]0;%d]", days[i], i, minDay, maxDay)
		}
	}
	return &byYearDayPart{
		partName: BySetPos,
		days:     types.NewIntegerValues(days),
	}, nil
}

func (p *byYearDayPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *byYearDayPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *byYearDayPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.days); i++ {
		secondsOutput.Write([]byte(p.days[i].GetStringValue()))
		if len(p.days)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
