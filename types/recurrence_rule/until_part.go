package recurrence_rule

import (
	"fmt"
	"io"
	"time"

	"github.com/vareversat/gics/types"
)

type UntilPart interface {
	RRPart
}

type untilPart struct {
	PartName RRPartName
	EndDate  types.DateTimeValue
}

func NewUntilPart(endDate time.Time) UntilPart {
	return &untilPart{
		PartName: UNTIL,
		EndDate:  types.NewDateTimeValue(endDate, types.WithLocalTime),
	}
}

func (f untilPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f untilPart) GetPartName() RRPartName {
	return f.PartName
}

func (f untilPart) GetPartValue() string {
	return f.EndDate.GetValue()
}
