package types

import (
	"reflect"
	"testing"
	"time"
)

func TestNewImplicitPeriodOfTime_GetStringValue(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"April 22th, 1997 at 8:54pm + 1d", args{
			from: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to:   86400 * 1e9}, "19970422T205400Z/P1D"},
		{"April 22th, 1997 at 8:54pm + 1h", args{
			from: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to:   3600 * 1e9}, "19970422T205400Z/PT1H"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			periodOfTime := NewImplicitPeriodOfTimeValue(tt.args.from, tt.args.to)
			if got := periodOfTime.GetStringValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("periodOfTime.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewExplicitPeriodOfTime_GetStringValue(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"April 22th, 1997 at 8:54pm + 1d", args{
			from: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to: time.Date(
				1997,
				4,
				23,
				20,
				54,
				0,
				0,
				time.UTC,
			)}, "19970422T205400Z/19970423T205400Z"},
		{"April 22th, 1997 at 8:54pm + 1h", args{
			from: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to: time.Date(
				1997,
				4,
				22,
				21,
				54,
				0,
				0,
				time.UTC,
			)}, "19970422T205400Z/19970422T215400Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			periodOfTime := NewExplicitPeriodOfTimeValue(tt.args.from, tt.args.to)
			if got := periodOfTime.GetStringValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("periodOfTime.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewImplicitPeriodOfTime_GetValue(t *testing.T) {
	type args struct {
		value string
	}
	type want struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"April 22th, 1997 at 8:54pm + 1d",
			args{value: "19970422T205400Z/P1D"},
			want{
				from: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
				to:   time.Date(1997, 4, 23, 20, 54, 0, 0, time.UTC)}},
		{"April 22th, 1997 at 8:54pm + 1h",
			args{value: "19970422T205400Z/PT1H"}, want{
				from: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
				to:   time.Date(1997, 4, 22, 21, 54, 0, 0, time.UTC)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			periodOfTime, _ := NewPeriodOfTimeValueFromString(tt.args.value)
			if from, _ := periodOfTime.GetValue(); !reflect.DeepEqual(from, tt.want.from) {
				t.Errorf("(from) periodOfTime.GetValue() = %v, want %v", from, tt.want.from)
			}
			if _, to := periodOfTime.GetValue(); !reflect.DeepEqual(to, tt.want.to) {
				t.Errorf("(to) periodOfTime.GetValue() = %v, want %v", to, tt.want.to)
			}
			if format := periodOfTime.GetFormat(); !reflect.DeepEqual(
				format,
				ImplicitPeriodOfTime,
			) {
				t.Errorf(
					"(format) periodOfTime.GetFormat() = %v, want %v",
					format,
					ImplicitPeriodOfTime,
				)
			}
		})
	}
}

func TestNewExplicitPeriodOfTime_GetValue(t *testing.T) {
	type args struct {
		value string
	}
	type want struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"April 22th, 1997 at 8:54pm + 1d", args{value: "19970422T205400Z/19970423T205400Z"}, want{
			from: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to:   time.Date(1997, 4, 23, 20, 54, 0, 0, time.UTC)}},
		{"April 22th, 1997 at 8:54pm + 1h", args{value: "19970422T205400Z/19970422T215400Z"}, want{
			from: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to:   time.Date(1997, 4, 22, 21, 54, 0, 0, time.UTC)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			periodOfTime, _ := NewPeriodOfTimeValueFromString(tt.args.value)
			if from, _ := periodOfTime.GetValue(); !reflect.DeepEqual(from, tt.want.from) {
				t.Errorf("(from) periodOfTime.GetValue() = %v, want %v", from, tt.want.from)
			}
			if _, to := periodOfTime.GetValue(); !reflect.DeepEqual(to, tt.want.to) {
				t.Errorf("(to) periodOfTime.GetValue() = %v, want %v", to, tt.want.to)
			}
			if format := periodOfTime.GetFormat(); !reflect.DeepEqual(
				format,
				ImplicitPeriodOfTime,
			) {
				t.Errorf(
					"(format) periodOfTime.GetFormat() = %v, want %v",
					format,
					ExplicitPeriodOfTime,
				)
			}
		})
	}
}
