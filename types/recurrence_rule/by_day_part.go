package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	minWeekOrdinal = -53
	maxWeekOrdinal = 53
)

type ByDayPart interface {
	RecurrenceRulePart
}

type byDayPart struct {
	partName RecurrenceRulePartName
	weekDays WeekDayWithOrdinals
}

type WeekDayWithOrdinal interface {
	ToString() string
}

type weekDayWithOrdinal struct {
	ordinal int32
	weekday WeekDay
}

type WeekDayWithOrdinals []WeekDayWithOrdinal

type WeekDay string

const (
	Monday    WeekDay = "MO"
	Tuesday   WeekDay = "TU"
	Wednesday WeekDay = "WE"
	Thursday  WeekDay = "TH"
	Friday    WeekDay = "FR"
	Saturday  WeekDay = "SA"
	Sunday    WeekDay = "SU"
)

// NewByDayPart give the info on which day of the month the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYDAY=1TU => "11th tuesday occurrence in the year"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByDayPart(days []WeekDayWithOrdinal) ByDayPart {
	return &byDayPart{
		partName: ByDay,
		weekDays: days,
	}
}

// NewByDayPart give the info on which day of the month the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYDAY=1TU => "11th tuesday occurrence in the year"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByDayPartFromString(value string) (ByDayPart, error) {
	var weekDays WeekDayWithOrdinals
	var err error
	stringWeekDays := strings.Split(value, multiValueSeparator)
	for _, weekDay := range stringWeekDays {
		var parsedItems int
		var finalOrdinal int32
		var finalWeekDay string
		parsedItems, err = fmt.Sscanf(strings.TrimSpace(weekDay), "%d%s", &finalOrdinal, &finalWeekDay)

		if err != nil || parsedItems != 2 {
			return nil, fmt.Errorf(
				"cannot parsed the following string into BYDAY component %s",
				err.Error(),
			)
		} else {
			weekDayWithOrdinal, err := NewWeekDayWithOrdinal(finalOrdinal, WeekDay(finalWeekDay))
			if err != nil {
				return nil, err
			} else {
				weekDays = append(weekDays, weekDayWithOrdinal)
			}
		}
	}
	return NewByDayPart(weekDays), nil

}

// NewWeekDayWithOrdinal return an ordinal followed by a week name. See [RFC-5545] ref for more info
// Example: 11TU => "11th tuesday occurrence in the year"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewWeekDayWithOrdinal(ordinal int32, weekDay WeekDay) (WeekDayWithOrdinal, error) {
	if ordinal < minWeekOrdinal || ordinal > maxWeekOrdinal || ordinal == 0 {
		return nil, fmt.Errorf(
			"%d is not a valid ordinal. It must satisfy this : ord ∈ [%d;0[U]0;%d]",
			ordinal,
			minWeekOrdinal,
			maxWeekOrdinal,
		)
	}
	return &weekDayWithOrdinal{
		ordinal: ordinal,
		weekday: weekDay,
	}, nil
}

func (p *byDayPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *byDayPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *byDayPart) GetPartValue() string {
	var output bytes.Buffer
	for i := 0; i < len(p.weekDays); i++ {
		output.Write([]byte(p.weekDays[i].ToString()))
		if len(p.weekDays)-1 > i {
			output.Write([]byte(","))
		}
	}
	return output.String()
}

func (w *weekDayWithOrdinal) ToString() string {
	return fmt.Sprintf("%d%s", w.ordinal, w.weekday)
}
