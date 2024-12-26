package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type BySecondPart interface {
	RecurrenceRulePart
}

type bySecondPart struct {
	PartName RecurrenceRulePartName
	Seconds  []types.IntegerType
}

// NewBySecondPart seconds in range [0,60]
func NewBySecondPart(seconds []int32) BySecondPart {
	return &bySecondPart{
		PartName: BYSECOND,
		Seconds:  types.NewIntegerValues(seconds),
	}
}

func (f bySecondPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f bySecondPart) GetPartName() RecurrenceRulePartName {
	return f.PartName
}

func (f bySecondPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.Seconds); i++ {
		secondsOutput.Write([]byte(f.Seconds[i].GetStringValue()))
		if len(f.Seconds)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}
