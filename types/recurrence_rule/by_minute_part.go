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
	minMinute = 0
	maxMinute = 59
)

type ByMinutePart interface {
	RecurrenceRulePart
}

type byMinutePart struct {
	partName RecurrenceRulePartName
	minutes  []types.IntegerType
}

// NewByMinutePart give the info on which minute of the day the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYMINUTE=19 => "at XX:19 (XX:19 am/pm)"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByMinutePart(minutes []int32) (ByMinutePart, error) {
	for i := 0; i < len(minutes); i++ {
		if minutes[i] < minMinute || minutes[i] > maxMinute {
			return nil, fmt.Errorf("%d (%d-th) is not a valid minute. It must satisfy this : minute ∈ [%d;%d]", minutes[i], i, minMinute, maxMinute)
		}
	}
	return &byMinutePart{
		partName: ByMinute,
		minutes:  types.NewIntegerValues(minutes),
	}, nil
}

// NewByMinutePart give the info on which minute of the day the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYMINUTE=19 => "at XX:19 (XX:19 am/pm)"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewByMinutePartFromString(value string) (ByMinutePart, error) {
	var minutes []int32
	stringMinutes := strings.Split(value, multiValueSeparator)
	for _, minute := range stringMinutes {
		finalMinute, err := strconv.ParseInt(strings.TrimSpace(minute), 10, 32)

		if err != nil {
			return nil, fmt.Errorf("cannot parsed the following string into int %s", err.Error())
		} else {
			minutes = append(minutes, int32(finalMinute))
		}
	}
	return NewByMinutePart(minutes)

}

func (p *byMinutePart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *byMinutePart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *byMinutePart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.minutes); i++ {
		secondsOutput.Write([]byte(p.minutes[i].GetStringValue()))
		if len(p.minutes)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
