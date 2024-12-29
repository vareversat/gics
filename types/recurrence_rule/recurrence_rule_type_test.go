package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewRecurrenceRuleValue(t *testing.T) {
	weekDay, _ := NewWeekDayWithOrdinal(5, Sunday)
	byMonthPart, _ := NewByMonthPart([]int32{12})
	type args struct {
		parts RecurrenceRuleParts
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Create a RRULE #1 - Every YEAR, on the 5th SUNDAY of the 12th MONTH",
			args{parts: RecurrenceRuleParts{
				NewFrequencyPart(Yearly),
				NewIntervalPart(1),
				byMonthPart,
				NewByDayPart(WeekDayWithOrdinals{weekDay}),
			}},
			"FREQ=YEARLY;INTERVAL=1;BYMONTH=12;BYDAY=5SU",
		},
		{"Create a RRULE #2 - Every 10 DAY, 10 times", args{parts: RecurrenceRuleParts{
			NewFrequencyPart(Daily),
			NewIntervalPart(10),
			NewCountPart(10),
		}}, "FREQ=DAILY;INTERVAL=10;COUNT=10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRecurrenceRuleValue(tt.args.parts); !reflect.DeepEqual(
				got.GetStringValue(),
				tt.want,
			) {
				t.Errorf("NewRecurrenceRuleValue() = %s, want %s", got.GetStringValue(), tt.want)
			}
		})
	}
}
