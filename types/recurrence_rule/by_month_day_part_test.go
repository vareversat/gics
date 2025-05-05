package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewByMonthDayPart(t *testing.T) {
	type args struct {
		monthDay []int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Create BYMONTHDAY Recurrence rule component #1",
			args{monthDay: []int32{10}},
			"10",
			false,
		},
		{
			"Create BYMONTHDAY Recurrence rule component #2",
			args{monthDay: []int32{10, 22}},
			"10,22",
			false,
		},
		{
			"Create BYMONTHDAY Recurrence rule component #3",
			args{monthDay: []int32{10, 22, 100}},
			"0",
			true,
		},
		{
			"Create BYMONTHDAY Recurrence rule component #4",
			args{monthDay: []int32{10, 22, -100}},
			"0",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByMonthDayPart(tt.args.monthDay)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByMonthDayPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}

func TestNewByMonthDayPartFromString(t *testing.T) {
	want, _ := NewByMonthDayPart([]int32{10, 11})
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    ByMonthDayPart
		wantErr bool
	}{
		{"Create ByMonth (no error)", args{value: "10,11"}, want, false},
		{"Create ByMonth (with error)", args{value: "100"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByMonthDayPartFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("got \"%s\" error, want not", err.Error())
			} else if got != nil && !reflect.DeepEqual(got.GetPartValue(), tt.want.GetPartValue()) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want.GetPartValue())
			}
		})
	}
}
