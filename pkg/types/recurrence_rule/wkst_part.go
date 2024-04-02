package recurrence_rule

import (
	"fmt"
	"io"
)

type WKSTPart interface {
	RRPart
}

type wKSTPart struct {
	PartName RRPartName
	WeekDay  RRWeekday
}

func NewWKSTPart(weekday RRWeekday) WKSTPart {
	return &wKSTPart{
		PartName: WKST,
		WeekDay:  weekday,
	}
}

func (f wKSTPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f wKSTPart) GetPartName() RRPartName {
	return f.PartName
}

func (f wKSTPart) GetPartValue() string {
	return string(f.WeekDay)
}
