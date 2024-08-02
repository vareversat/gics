package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.3

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types/recurrence_rule"
)

type RecurrenceRuleProperty interface {
	RequestStatusPropertyType
}

func NewRecurrenceRuleProperty(parts ...recurrence_rule.RRPart) RecurrenceRuleProperty {
	return &recurrenceRulePropertyType{
		PropName: registries.RecurrenceRuleProp,
		Value:    recurrence_rule.NewRecurrenceRuleValue(parts),
	}
}
