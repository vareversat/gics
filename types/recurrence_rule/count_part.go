package recurrence_rule

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/vareversat/gics/types"
)

type CountPart interface {
	RecurrenceRulePart
}

type countPart struct {
	partName RecurrenceRulePartName
	count    types.IntegerType
}

// NewCountPart give the info on how many time repeat the recurrence rule. See [RFC-5545] ref for more info
// Example: COUNT=100 => "repeat 100 times"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewCountPart(count int32) CountPart {
	return &countPart{
		partName: Count,
		count:    types.NewIntegerValue(count),
	}
}

// NewCountPart give the info on how many time repeat the recurrence rule. See [RFC-5545] ref for more info
// Example: COUNT=100 => "repeat 100 times"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewCountPartFromString(value string) (CountPart, error) {
	count, err := strconv.ParseInt(strings.TrimSpace(value), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("cannot parsed the following string into int %s", err.Error())
	}
	return NewCountPart(int32(count)), nil

}

func (p *countPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *countPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *countPart) GetPartValue() string {
	return p.count.GetStringValue()
}
