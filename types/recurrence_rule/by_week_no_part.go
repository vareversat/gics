package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type ByWeekNoPart interface {
	RRPart
}

type byWeekNoPart struct {
	PartName RRPartName
	Seconds  []types.IntegerType
}

// NewByWeekNoPart days in range [-53,0[U]0,53]
func NewByWeekNoPart(weeks []int32) ByWeekNoPart {
	return &byWeekNoPart{
		PartName: BYWEEKNO,
		Seconds:  types.NewIntegerValues(weeks),
	}
}

func (f byWeekNoPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f byWeekNoPart) GetPartName() RRPartName {
	return f.PartName
}

func (f byWeekNoPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.Seconds); i++ {
		secondsOutput.Write([]byte(f.Seconds[i].GetStringValue()))
		if len(f.Seconds)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}
