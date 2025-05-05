package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/vareversat/gics/types"
)

const (
	minMonth = 1
	maxMonth = 13
)

type ByMonthPart interface {
	RecurrenceRulePart
}

type byMonthPart struct {
	partName RecurrenceRulePartName
	months   []types.IntegerType
}

// NewByMonthPart give the info on which month of the year the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYMONTH=2 => "2nd month of the year"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByMonthPart(months []int32) (ByMonthPart, error) {
	for i := 0; i < len(months); i++ {
		if months[i] < minMonth || months[i] > maxMonth {
			return nil, fmt.Errorf("%d (%d-th) is not a valid month. It must satisfy this : month ∈ [%d;%d]", months[i], i, minMonth, maxMonth)
		}
	}
	return &byMonthPart{
		partName: ByMonth,
		months:   types.NewIntegerValues(months),
	}, nil
}

// NewByMonthPartFromString give the info on which month of the year the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYMONTH=2 => "2nd month of the year"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByMonthPartFromString(value string) (ByMonthPart, error) {
	var months []int32
	stringMonths := strings.Split(value, multiValueSeparator)
	for _, month := range stringMonths {
		finalMonth, err := strconv.ParseInt(strings.TrimSpace(month), 10, 32)

		if err != nil {
			return nil, fmt.Errorf("cannot parsed the following string into int %s", err.Error())
		} else {
			months = append(months, int32(finalMonth))
		}
	}
	return NewByMonthPart(months)

}

func (p *byMonthPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *byMonthPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *byMonthPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.months); i++ {
		secondsOutput.Write([]byte(p.months[i].GetStringValue()))
		if len(p.months)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
