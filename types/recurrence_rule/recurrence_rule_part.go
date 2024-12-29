package recurrence_rule

import "io"

type RecurrenceRulePartName string

const (
	Frequency    RecurrenceRulePartName = "FREQ"
	Until        RecurrenceRulePartName = "UNTIL"
	Count        RecurrenceRulePartName = "COUNT"
	Interval     RecurrenceRulePartName = "INTERVAL"
	BySecond     RecurrenceRulePartName = "BYSECOND"
	ByMinute     RecurrenceRulePartName = "BYMINUTE"
	ByHour       RecurrenceRulePartName = "BYHOUR"
	ByDay        RecurrenceRulePartName = "BYDAY"
	ByMonthDay   RecurrenceRulePartName = "BYMONTHDAY"
	ByYearDay    RecurrenceRulePartName = "BYYEARDAY"
	ByWeekNumber RecurrenceRulePartName = "BYWEEKNO"
	ByMonth      RecurrenceRulePartName = "BYMONTH"
	BySetPos     RecurrenceRulePartName = "BYSETPOS"
	WeekStart    RecurrenceRulePartName = "WKST"
)

type RecurrenceRulePart interface {
	// GetPartName return the name of the Recurrence rule part
	GetPartName() RecurrenceRulePartName
	// GetPartValue return the string value of the Recurrence rule part
	GetPartValue() string
	// ToICalendarPartFormat return the formatted the Recurrence rule part ('NAME=VALUE')
	ToICalendarPartFormat(output io.Writer)
}

type RecurrenceRuleParts []RecurrenceRulePart
