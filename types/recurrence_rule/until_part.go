package recurrence_rule

import (
	"fmt"
	"io"
	"time"

	"github.com/vareversat/gics/types"
)

type UntilPart interface {
	RecurrenceRulePart
}

type untilPart struct {
	PartName RecurrenceRulePartName
	EndDate  types.DateTimeType
}

func NewUntilPart(endDate time.Time) UntilPart {
	return &untilPart{
		PartName: UNTIL,
		EndDate:  types.NewDateTimeValue(endDate),
	}
}

func (f untilPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f untilPart) GetPartName() RecurrenceRulePartName {
	return f.PartName
}

func (f untilPart) GetPartValue() string {
	return f.EndDate.GetStringValue()
}
