package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types/recurrence_rule"
)

type RecurrenceRuleProperty interface {
	RequestStatusPropertyType
}

func NewRecurrenceRuleProperty(parts ...recurrence_rule.RRPart) RecurrenceRuleProperty {
	return &recurrenceRulePropertyType{
		PropName: registries.RRULE,
		Value:    recurrence_rule.NewRecurrenceRuleValue(parts),
	}
}
