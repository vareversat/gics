package properties

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types/recurrence_rule"
)

type ExceptionRuleProperty interface {
	RequestStatusPropertyType
}

// NewExceptionRuleProperty create a new registries.ExceptionRuleProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc2445#section-4.8.5.2
// NewExceptionRuleProperty Deprecated
func NewExceptionRuleProperty(parts ...recurrence_rule.RRPart) ExceptionRuleProperty {
	return &recurrenceRulePropertyType{
		PropName: registries.ExceptionRuleProp,
		Value:    recurrence_rule.NewRecurrenceRuleValue(parts),
	}
}
