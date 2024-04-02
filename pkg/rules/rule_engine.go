package rules

import "github.com/vareversat/gics/pkg/components"

type RuleEngine interface {
}

type ruleEngine struct {
}

func (e *ruleEngine) hasAllMandatoryProperties(component components.CalendarComponent) {

}
