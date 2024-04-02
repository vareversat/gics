package recurrence_rule

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/pkg/types"
)

type IntervalPart interface {
	RRPart
}

type intervalPart struct {
	PartName RRPartName
	Count    types.IntegerValue
}

// NewIntervalPart count in range [0,9]
func NewIntervalPart(count int32) IntervalPart {
	return &intervalPart{
		PartName: INTERVAL,
		Count:    types.NewIntegerValue(count),
	}
}

func (f intervalPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f intervalPart) GetPartName() RRPartName {
	return f.PartName
}

func (f intervalPart) GetPartValue() string {
	return f.Count.GetStringValue()
}
