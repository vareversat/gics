package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/vareversat/gics/types"
)

type BySetPosPart interface {
	RecurrenceRulePart
}

type bySetPosPart struct {
	partName    RecurrenceRulePartName
	setPosition []types.IntegerType
}

// NewBySetPosPart give the info on which occurrence within the set of recurrence instances specified by the rule. See [RFC-5545] ref for more info
// Example: FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-1 => "the last work day of the month"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewBySetPosPart(setPosition []int32) (BySetPosPart, error) {
	for i := 0; i < len(setPosition); i++ {
		if setPosition[i] < minDay || setPosition[i] > maxDay || setPosition[i] == 0 {
			return nil, fmt.Errorf("%d (%d-th) is not a valid set position. It must satisfy this : setpos ∈ [%d;0[U]0;%d]", setPosition[i], i, minDay, maxDay)
		}
	}
	return &bySetPosPart{
		partName:    BySetPos,
		setPosition: types.NewIntegerValues(setPosition),
	}, nil
}

// NewBySetPosPartFromString give the info on which occurrence within the set of recurrence instances specified by the rule. See [RFC-5545] ref for more info
// Example: FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-1 => "the last work day of the month"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewBySetPosPartFromString(value string) (BySetPosPart, error) {
	var positions []int32
	stringPositions := strings.Split(value, multiValueSeparator)
	for _, position := range stringPositions {
		finalPosition, err := strconv.ParseInt(strings.TrimSpace(position), 10, 32)

		if err != nil {
			return nil, fmt.Errorf("cannot parsed the following string into int %s", err.Error())
		} else {
			positions = append(positions, int32(finalPosition))
		}
	}
	return NewBySetPosPart(positions)

}

func (p *bySetPosPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *bySetPosPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *bySetPosPart) GetPartValue() string {
	var positionsOutput bytes.Buffer
	for i := 0; i < len(p.setPosition); i++ {
		positionsOutput.Write([]byte(p.setPosition[i].GetStringValue()))
		if len(p.setPosition)-1 > i {
			positionsOutput.Write([]byte(","))
		}
	}
	return positionsOutput.String()
}
