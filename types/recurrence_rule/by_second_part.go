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
	minSecond = 0
	maxSecond = 60
)

type BySecondPart interface {
	RecurrenceRulePart
}

type bySecondPart struct {
	partName RecurrenceRulePartName
	seconds  []types.IntegerType
}

// NewBySecondPart give the info on which second of the minute the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYSECOND=2 => "2nd second of the minute"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewBySecondPart(seconds []int32) (BySecondPart, error) {
	for i := 0; i < len(seconds); i++ {
		if seconds[i] < minSecond || seconds[i] > maxSecond {
			return nil, fmt.Errorf("%d (%d-th) is not a valid second. It must satisfy this : second ∈ [%d;%d]", seconds[i], i, minSecond, maxSecond)
		}
	}
	return &bySecondPart{
		partName: BySecond,
		seconds:  types.NewIntegerValues(seconds),
	}, nil
}

// NewBySecondPartFromString give the info on which second of the minute the recurrence occurs. See [RFC-5545] ref for more info
// Example: BYSECOND=2 => "2nd second of the minute"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewBySecondPartFromString(value string) (BySecondPart, error) {
	var seconds []int32
	stringSeconds := strings.Split(value, multiValueSeparator)
	for _, second := range stringSeconds {
		finalSecond, err := strconv.ParseInt(strings.TrimSpace(second), 10, 32)

		if err != nil {
			return nil, fmt.Errorf("cannot parsed the following string into int %s", err.Error())
		} else {
			seconds = append(seconds, int32(finalSecond))
		}
	}
	return NewBySecondPart(seconds)

}

func (p *bySecondPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *bySecondPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *bySecondPart) GetPartValue() string {
	var secondsOutput bytes.Buffer
	for i := 0; i < len(p.seconds); i++ {
		secondsOutput.Write([]byte(p.seconds[i].GetStringValue()))
		if len(p.seconds)-1 > i {
			secondsOutput.Write([]byte(","))
		}
	}
	return secondsOutput.String()
}
