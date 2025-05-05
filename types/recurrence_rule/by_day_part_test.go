package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewByDayPart(t *testing.T) {
	week1, _ := NewWeekDayWithOrdinal(1, Tuesday)
	week2, _ := NewWeekDayWithOrdinal(20, Tuesday)
	type args struct {
		days WeekDayWithOrdinals
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Create BYDAY Recurrence rule component #1",
			args{days: WeekDayWithOrdinals{week1}},
			"1TU",
		},
		{
			"Create BYDAY Recurrence rule component #2",
			args{days: WeekDayWithOrdinals{week2}},
			"20TU",
		},
		{
			"Create BYDAY Recurrence rule component #3",
			args{days: WeekDayWithOrdinals{week1, week2}},
			"1TU,20TU",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewByDayPart(tt.args.days); !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.GetPartValue() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}

func TestNewWeekDayWithOrdinal(t *testing.T) {
	type args struct {
		ordinal int32
		weekDay WeekDay
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Create WeekDay (no error)", args{ordinal: 1, weekDay: Tuesday}, "1TU", false},
		{"Create WeekDay (with error)", args{ordinal: 100, weekDay: Tuesday}, "0SU", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWeekDayWithOrdinal(tt.args.ordinal, tt.args.weekDay)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWeekDayWithOrdinal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ToString(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.ToString(), tt.want)
			}
		})
	}
}

func TestNewWeekDayWithOrdinalFromString(t *testing.T) {
	var weekDays []WeekDayWithOrdinal
	wdo1, _ := NewWeekDayWithOrdinal(1, Tuesday)
	wdo2, _ := NewWeekDayWithOrdinal(2, Wednesday)
	weekDays = append(weekDays, wdo1, wdo2)
	want := NewByDayPart(weekDays)
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    ByDayPart
		wantErr bool
	}{
		{
			"Create WeekDay (no error)",
			args{value: "1TU, 2WE"},
			want,
			false,
		},
		{
			"Create WeekDay (with error)",
			args{value: "1000000MO"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByDayPartFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("got \"%s\" error, want not", err.Error())
			} else if got != nil && !reflect.DeepEqual(got.GetPartValue(), tt.want.GetPartValue()) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want.GetPartValue())
			}
		})
	}
}
