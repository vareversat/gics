package recurrence_rule

import (
	"fmt"
	"io"
)

type RecurrenceRuleFrequency string

const (
	Secondly RecurrenceRuleFrequency = "SECONDLY"
	Minutely RecurrenceRuleFrequency = "MINUTELY"
	Hourly   RecurrenceRuleFrequency = "HOURLY"
	Daily    RecurrenceRuleFrequency = "DAILY"
	Weekly   RecurrenceRuleFrequency = "WEEKLY"
	Monthly  RecurrenceRuleFrequency = "MONTHLY"
	Yearly   RecurrenceRuleFrequency = "YEARLY"
)

type FrequencyPart interface {
	RecurrenceRulePart
}

type frequencyPart struct {
	partName  RecurrenceRulePartName
	frequency RecurrenceRuleFrequency
}

// NewFrequencyPart give the info of the recurrence occurs type. See [RFC-5545] ref for more info
// Example: FREQ=SECONDLY => "events based on an interval of a second or more"
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewFrequencyPart(frequency RecurrenceRuleFrequency) FrequencyPart {
	return &frequencyPart{
		partName:  Frequency,
		frequency: frequency,
	}
}

func (p *frequencyPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", p.GetPartName(), p.GetPartValue())))
}

func (p *frequencyPart) GetPartName() RecurrenceRulePartName {
	return p.partName
}

func (p *frequencyPart) GetPartValue() string {
	return string(p.frequency)
}
