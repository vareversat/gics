package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type ByYearDayPart interface {
	RecurrenceRulePart
}

type byYearDayPart struct {
	PartName RecurrenceRulePartName
	Seconds  []types.IntegerType
}

// NewByYearDayPart days in range [-366,0[U]0,366]
func NewByYearDayPart(days []int32) ByYearDayPart {
	return &byYearDayPart{
		PartName: BYYEARDAY,
		Seconds:  types.NewIntegerValues(days),
	}
}

func (f byYearDayPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f byYearDayPart) GetPartName() RecurrenceRulePartName {
	return f.PartName
}

func (f byYearDayPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.Seconds); i++ {
		secondsOutput.Write([]byte(f.Seconds[i].GetStringValue()))
		if len(f.Seconds)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}
