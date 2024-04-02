package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/pkg/types"
)

type BySetPosPart interface {
	RRPart
}

type bySetPosPart struct {
	PartName RRPartName
	Seconds  []types.IntegerValue
}

// NewBySetPosPart days in range [-366,0[U]0,366]
func NewBySetPosPart(days []int32) BySetPosPart {
	return &bySetPosPart{
		PartName: BYSETPOS,
		Seconds:  types.NewIntegerValues(days),
	}
}

func (f bySetPosPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f bySetPosPart) GetPartName() RRPartName {
	return f.PartName
}

func (f bySetPosPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.Seconds); i++ {
		secondsOutput.Write([]byte(f.Seconds[i].GetStringValue()))
		if len(f.Seconds)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}
