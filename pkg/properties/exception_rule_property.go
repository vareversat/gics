package properties

// https://datatracker.ietf.org/doc/html/rfc2445#section-4.8.5.2

// DEPRECATED

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types/recurrence_rule"
)

type ExceptionRuleProperty interface {
	RequestStatusPropertyType
}

func NewExceptionRuleProperty(parts ...recurrence_rule.RRPart) ExceptionRuleProperty {
	return &recurrenceRulePropertyType{
		PropName: registries.EXRULE,
		Value:    recurrence_rule.NewRecurrenceRuleValue(parts),
	}
}
