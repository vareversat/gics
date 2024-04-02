package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/pkg/types"
)

type ByMonthDayPart interface {
	RRPart
}

type byMonthDayPart struct {
	PartName RRPartName
	Seconds  []types.IntegerValue
}

// NewByMonthDayPart months in range [-31,0[U]0,31]
func NewByMonthDayPart(months []int32) ByMonthDayPart {
	return &byMonthDayPart{
		PartName: BYMONTHDAY,
		Seconds:  types.NewIntegerValues(months),
	}
}

func (f byMonthDayPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f byMonthDayPart) GetPartName() RRPartName {
	return f.PartName
}

func (f byMonthDayPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.Seconds); i++ {
		secondsOutput.Write([]byte(f.Seconds[i].GetStringValue()))
		if len(f.Seconds)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}
