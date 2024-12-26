package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type ByMonthPart interface {
	RecurrenceRulePart
}

type byMonthPart struct {
	PartName RecurrenceRulePartName
	Seconds  []types.IntegerType
}

// NewByMonthPart days in range [1,13]
func NewByMonthPart(weeks []int32) ByMonthPart {
	return &byMonthPart{
		PartName: BYMONTH,
		Seconds:  types.NewIntegerValues(weeks),
	}
}

func (f byMonthPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f byMonthPart) GetPartName() RecurrenceRulePartName {
	return f.PartName
}

func (f byMonthPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.Seconds); i++ {
		secondsOutput.Write([]byte(f.Seconds[i].GetStringValue()))
		if len(f.Seconds)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}
