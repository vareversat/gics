package properties

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types/recurrence_rule"
)

type RecurrenceRuleProperty interface {
	RequestStatusPropertyType
}

// NewRecurrenceRuleProperty create a new registries.RecurrenceRuleProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional if recurring)
// - registries.Vtodo (Optional if recurring)
// - registries.Vjournal (Optional if recurring)
// - registries.Standard (Optional if recurring)
// - registries.Daylight (Optional if recurring)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// - registries.RecurrenceIdRangeParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.3
func NewRecurrenceRuleProperty(parts ...recurrence_rule.RecurrenceRulePart) RecurrenceRuleProperty {
	return &recurrenceRulePropertyType{
		PropName: registries.RecurrenceRuleProp,
		Value:    recurrence_rule.NewRecurrenceRuleValue(parts),
	}
}
