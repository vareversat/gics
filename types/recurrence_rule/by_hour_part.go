package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/types"
)

const (
	minHour = 0
	maxHour = 23
)

type ByHourPart interface {
	RecurrenceRulePart
}

type byHourPart struct {
	partName RecurrenceRulePartName
	hours    []types.IntegerType
}

// NewByHourPart give the info on which day of the month the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYHOUR=19 => "at 19:XX (7:XX pm)"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByHourPart(hours []int32) (ByHourPart, error) {
	for i := 0; i < len(hours); i++ {
		if hours[i] < minHour || hours[i] > maxHour {
			return &byHourPart{
				partName: ByHour,
				hours:    types.NewIntegerValues([]int32{0}),
			}, fmt.Errorf("%d (%d-th) is not a valid hour. It must satisfy this : hour ∈ [%d;%d]", hours[i], i, minHour, maxHour)
		}
	}
	return &byHourPart{
		partName: ByHour,
		hours:    types.NewIntegerValues(hours),
	}, nil
}

func (p *byHourPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *byHourPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *byHourPart) GetPartValue() string {
	var output bytes.Buffer
	for i := 0; i < len(p.hours); i++ {
		output.Write([]byte(p.hours[i].GetStringValue()))
		if len(p.hours)-1 > i {
			output.Write([]byte((",")))
		}
	}
	return output.String()
}
