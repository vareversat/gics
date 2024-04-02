package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10

// Example :
// - FREQ=DAILY;COUNT=10
// - FREQ=DAILY;UNTIL=19971224T000000Z
// - FREQ=DAILY;UNTIL=20000131T140000Z;BYMONTH=1

type RRPartName string

const (
	FREQ       RRPartName = "FREQ"
	UNTIL      RRPartName = "UNTIL"
	COUNT      RRPartName = "COUNT"
	INTERVAL   RRPartName = "INTERVAL"
	BYSECOND   RRPartName = "BYSECOND"
	BYMINUTE   RRPartName = "BYMINUTE"
	BYHOUR     RRPartName = "BYHOUR"
	BYDAY      RRPartName = "BYDAY"
	BYMONTHDAY RRPartName = "BYMONTHDAY"
	BYYEARDAY  RRPartName = "BYYEARDAY"
	BYWEEKNO   RRPartName = "BYWEEKNO"
	BYMONTH    RRPartName = "BYMONTH"
	BYSETPOS   RRPartName = "BYSETPOS"
	WKST       RRPartName = "WKST"
)

type RRPart interface {
	ToICalendarPartFormat(output io.Writer)
	GetPartName() RRPartName
	GetPartValue() string
}

type RRParts []RRPart

type RecurrenceRuleValue struct {
	types.V
	Parts RRParts
}

func NewRecurrenceRuleValue(parts []RRPart) RecurrenceRuleValue {
	return RecurrenceRuleValue{
		V:     types.NewValue(registries.RECUR),
		Parts: parts,
	}
}

func (tV *RecurrenceRuleValue) GetValue() string {
	var partOut bytes.Buffer
	for i := 0; i < len(tV.Parts); i++ {
		if len(tV.Parts) > 1 && i > 0 {
			partOut.Write([]byte(fmt.Sprint(";")))
		}
		tV.Parts[i].ToICalendarPartFormat(&partOut)
	}
	return partOut.String()
}
