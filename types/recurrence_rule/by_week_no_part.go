package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type ByWeekNoPart interface {
	RecurrenceRulePart
}

type byWeekNoPart struct {
	partName   RecurrenceRulePartName
	weekNumber []types.IntegerType
}

// NewByWeekNoPart give the info on which week of the year the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYWEEKNO=10 => "10th week of the year"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByWeekNoPart(weekNumbers []int32) (ByWeekNoPart, error) {
	for i := 0; i < len(weekNumbers); i++ {
		if weekNumbers[i] < minWeekOrdinal || weekNumbers[i] > maxWeekOrdinal ||
			weekNumbers[i] == 0 {
			return &byWeekNoPart{
				partName:   ByWeekNumber,
				weekNumber: types.NewIntegerValues([]int32{0}),
			}, fmt.Errorf("%d (%d-th) is not a valid week number. It must satisfy this : ordinal ∈ [%d;0[U]0;%d]", weekNumbers[i], i, minWeekOrdinal, maxWeekOrdinal)
		}
	}
	return &byWeekNoPart{
		partName:   ByWeekNumber,
		weekNumber: types.NewIntegerValues(weekNumbers),
	}, nil
}

func (p *byWeekNoPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *byWeekNoPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *byWeekNoPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.weekNumber); i++ {
		secondsOutput.Write([]byte(p.weekNumber[i].GetStringValue()))
		if len(p.weekNumber)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
