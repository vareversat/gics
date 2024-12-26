package recurrence_rule

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type IntervalPart interface {
	RecurrenceRulePart
}

type intervalPart struct {
	PartName RecurrenceRulePartName
	Count    types.IntegerType
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

func (f intervalPart) GetPartName() RecurrenceRulePartName {
	return f.PartName
}

func (f intervalPart) GetPartValue() string {
	return f.Count.GetStringValue()
}
