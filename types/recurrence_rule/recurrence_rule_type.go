package recurrence_rule

import (
	"bytes"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RecurrenceRuleType interface {
	types.ValueType
}

type recurrenceRuleType struct {
	typeName  registries.ValueTypeRegistry
	typeValue RecurrenceRuleParts
}

// NewRecurrenceRuleValue create a new [registries.Recurrence] type. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
func NewRecurrenceRuleValue(parts RecurrenceRuleParts) RecurrenceRuleType {
	return &recurrenceRuleType{
		typeName:  registries.Recurrence,
		typeValue: parts,
	}
}

func (r *recurrenceRuleType) GetStringValue() string {
	var partOut bytes.Buffer
	for i := 0; i < len(r.typeValue); i++ {
		if len(r.typeValue) > 1 && i > 0 {
			partOut.Write([]byte(";"))
		}
		partOut.WriteString(r.typeValue[i].GetPartValue())
	}
	return partOut.String()
}

func (c *recurrenceRuleType) GetTypeName() string {
	return string(c.typeName)
}
