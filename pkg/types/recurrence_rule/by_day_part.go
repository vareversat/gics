package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"
)

type ByDayPart interface {
	RRPart
}

type byDayPart struct {
	PartName RRPartName
	WeekDays []RRWeekDayNum
}

func NewByDayPart(days []RRWeekDayNum) ByDayPart {
	return &byDayPart{
		PartName: BYDAY,
		WeekDays: days,
	}
}

func (f byDayPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f byDayPart) GetPartName() RRPartName {
	return f.PartName
}

func (f byDayPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(f.WeekDays); i++ {
		secondsOutput.Write([]byte(f.WeekDays[i].GetValue()))
		if len(f.WeekDays)-1 > i {
			secondsOutput.Write([]byte(fmt.Sprint(",")))
		}
	}
	return secondsOutput.String()
}

type RRWeekday string

const (
	SU RRWeekday = "SU"
	MO RRWeekday = "MO"
	TU RRWeekday = "MO"
	WE RRWeekday = "WE"
	TH RRWeekday = "TH"
	FR RRWeekday = "FR"
	SA RRWeekday = "SA"
)

// RRWeekDayNum
// Ordinal in range [0,23]
// Sign + or -
type RRWeekDayNum struct {
	Sign    string
	Ordinal int
	Weekday RRWeekday
}

func (rrdN *RRWeekDayNum) GetValue() string {
	return fmt.Sprintf("%s%d%s", rrdN.Sign, rrdN.Ordinal, rrdN.Weekday)
}
