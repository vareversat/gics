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
	partName RecurrenceRulePartName
	endDate  types.DateTimeType
}

func NewUntilPart(endDate time.Time) UntilPart {
	return &untilPart{
		partName: Until,
		endDate:  types.NewDateTimeValue(endDate),
	}
}

func (p *untilPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *untilPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *untilPart) GetPartValue() string {
	return p.endDate.GetStringValue()
}
