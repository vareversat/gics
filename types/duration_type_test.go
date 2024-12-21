package types

import (
	"reflect"
	"testing"
	"time"
)

func TestDurationGetStringValue(t *testing.T) {
	type args struct {
		value time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"19 years, 1 month, 4 weeks, 1 day, 20 hours, 14 minutes and 55 seconds",
			args{value: 604800898 * 1e9},
			"P19Y1M4W1DT20H14M55S",
		},
		{"1 week", args{value: 604800 * 1e9}, "P1W"},
		{"4 days", args{value: 345600 * 1e9}, "P4D"},
		{"1 week and 2 days", args{value: 777600 * 1e9}, "P1W2D"},
		{"2 week, 6 days and 11 hours", args{value: 1767600 * 1e9}, "P2W6DT11H"},
		{"2 week, 6 days, 11 hours and 16 seconds", args{value: 1767616 * 1e9}, "P2W6DT11H16S"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration := NewDurationValue(tt.args.value)
			if got := duration.GetStringValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("duration.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			"19 years, 1 month, 4 weeks, 1 day, 20 hours, 14 minutes and 55 seconds",
			args{value: "P19Y1M4W1DT20H14M55S"},
			604800898 * 1e9,
		},
		{"1 week", args{value: "P1W"}, 604800 * 1e9},
		{"4 days", args{value: "P4D"}, 345600 * 1e9},
		{"1 week and 2 days", args{value: "P1W2D"}, 777600 * 1e9},
		{"2 week, 6 days and 11 hours", args{value: "P2W6DT11H"}, 1767600 * 1e9},
		{"2 week, 6 days, 11 hours and 16 seconds", args{value: "P2W6DT11H0M16S"}, 1767616 * 1e9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration := NewDurationValueFromString(tt.args.value)
			if got := duration.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("duration.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
