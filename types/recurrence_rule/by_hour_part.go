package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type ByHourPart interface {
	RRPart
}

type byHourPart struct {
	PartName RRPartName
	Seconds  []types.IntegerType
}

// NewByHourPart seconds in range [0,23]
func NewByHourPart(seconds []int32) ByHourPart {
	return &byHourPart{
		PartName: BYHOUR,
		Seconds:  types.NewIntegerValues(seconds),
	}
}

func (f byHourPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f byHourPart) GetPartName() RRPartName {
	return f.PartName
}

func (f byHourPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.Seconds); i++ {
		secondsOutput.Write([]byte(f.Seconds[i].GetStringValue()))
		if len(f.Seconds)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}
