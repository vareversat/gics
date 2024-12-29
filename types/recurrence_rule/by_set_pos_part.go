package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type BySetPosPart interface {
	RecurrenceRulePart
}

type bySetPosPart struct {
	partName    RecurrenceRulePartName
	setPosition []types.IntegerType
}

// NewBySecondPart give the info on which occurrence within the set of recurrence instances specified by the rule. See [RFC-5545] ref for more info
// Example: FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-1 => "the last work day of the month"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewBySetPosPart(setPosition []int32) (BySetPosPart, error) {
	for i := 0; i < len(setPosition); i++ {
		if setPosition[i] < minDay || setPosition[i] > maxDay || setPosition[i] == 0 {
			return &bySetPosPart{
				partName:    BySetPos,
				setPosition: types.NewIntegerValues([]int32{0}),
			}, fmt.Errorf("%d (%d-th) is not a valid set position. It must satisfy this : setpos ∈ [%d;0[U]0;%d]", setPosition[i], i, minDay, maxDay)
		}
	}
	return &bySetPosPart{
		partName:    BySetPos,
		setPosition: types.NewIntegerValues(setPosition),
	}, nil
}

func (p *bySetPosPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *bySetPosPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *bySetPosPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.setPosition); i++ {
		secondsOutput.Write([]byte(p.setPosition[i].GetStringValue()))
		if len(p.setPosition)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
