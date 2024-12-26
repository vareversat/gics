package recurrence_rule

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

type CountPart interface {
	RecurrenceRulePart
}

type countPart struct {
	PartName RecurrenceRulePartName
	Count    types.IntegerType
}

// NewCountPart count in range [0,9]
func NewCountPart(count int32) CountPart {
	return &countPart{
		PartName: COUNT,
		Count:    types.NewIntegerValue(count),
	}
}

func (f countPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f countPart) GetPartName() RecurrenceRulePartName {
	return f.PartName
}

func (f countPart) GetPartValue() string {
	return f.Count.GetStringValue()
}
