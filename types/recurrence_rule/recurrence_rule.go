package recurrence_rule

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
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

type RecurrenceRuleType interface {
	types.ValueType
	GetValue() RRParts
}

type recurrenceRuleType struct {
	typeName  registries.ValueTypeRegistry
	typeValue RRParts
}

func NewRecurrenceRuleValue(parts []RRPart) RecurrenceRuleType {
	return &recurrenceRuleType{
		typeName:  registries.Recur,
		typeValue: parts,
	}
}

func (r *recurrenceRuleType) GetStringValue() string {
	var partOut bytes.Buffer
	for i := 0; i < len(r.typeValue); i++ {
		if len(r.typeValue) > 1 && i > 0 {
			partOut.Write([]byte(fmt.Sprint(";")))
		}
		r.typeValue[i].ToICalendarPartFormat(&partOut)
	}
	return partOut.String()
}

func (c *recurrenceRuleType) GetTypeName() string {
	return string(c.typeName)
}

// GetValue get the [registries.ComponentRegistry] typed value
func (tV *recurrenceRuleType) GetValue() RRParts {
	return nil
}
