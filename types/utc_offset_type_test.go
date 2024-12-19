package types

import (
	"reflect"
	"testing"
)

func TestNewUtcOffsetValue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1h", args{value: 3600}, "+0100"},
		{"5h30", args{value: 19800}, "+0530"},
		{"-11h", args{value: -39600}, "-1100"},
		{"-5h", args{value: -18000}, "-0500"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utcOffset := NewUtcOffsetValue(tt.args.value)
			if got := utcOffset.GetStringValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("utcOffset.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUtcOffsetValueFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1h", args{value: "+0100"}, 3600},
		{"5h30", args{value: "+0530"}, 19800},
		{"-11h", args{value: "-1100"}, -39600},
		{"-5h", args{value: "-0500"}, -18000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utcOffset := NewUtcOffsetValueFromString(tt.args.value)
			if got := utcOffset.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("utcOffset.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
