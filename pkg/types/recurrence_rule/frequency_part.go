package recurrence_rule

import (
	"fmt"
	"io"
)

type FrequencyPart interface {
	RRPart
}

type frequencyPart struct {
	PartName  RRPartName
	Frequency RRFrequency
}

func NewFrequencyPart(frequency RRFrequency) FrequencyPart {
	return &frequencyPart{
		PartName:  FREQ,
		Frequency: frequency,
	}
}

func (f frequencyPart) ToICalendarPartFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", f.GetPartName(), f.GetPartValue())))
}

func (f frequencyPart) GetPartName() RRPartName {
	return f.PartName
}

func (f frequencyPart) GetPartValue() string {
	return string(f.Frequency)
}

type RRFrequency string

const (
	SECONDLY RRFrequency = "SECONDLY"
	MINUTELY RRFrequency = "MINUTELY"
	HOURLY   RRFrequency = "HOURLY"
	DAILY    RRFrequency = "DAILY"
	WEEKLY   RRFrequency = "WEEKLY"
	MONTHLY  RRFrequency = "MONTHLY"
	YEARLY   RRFrequency = "YEARLY"
)
