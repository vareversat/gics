package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type ByMinutePart interface {
	RecurrenceRulePart
}

type byMinutePart struct {
	PartName RecurrenceRulePartName
	Seconds  []types.IntegerType
}

// NewByMinutePart seconds in range [0,59]
func NewByMinutePart(seconds []int32) ByMinutePart {
	return &byMinutePart{
		PartName: BYMINUTE,
		Seconds:  types.NewIntegerValues(seconds),
	}
}

func (f byMinutePart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f byMinutePart) GetPartName() RecurrenceRulePartName {
	return f.PartName
}

func (f byMinutePart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.Seconds); i++ {
		secondsOutput.Write([]byte(f.Seconds[i].GetStringValue()))
		if len(f.Seconds)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}
