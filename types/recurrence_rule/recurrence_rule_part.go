package recurrence_rule

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
	WKST         RecurrenceRulePartName = "WKST"
)

type RecurrenceRulePart interface {
	GetPartName() RecurrenceRulePartName
	GetPartValue() string
}

type RecurrenceRuleParts []RecurrenceRulePart
