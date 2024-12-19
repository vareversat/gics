package types

import (
	"reflect"
	"testing"
	"time"
)

func TestNewPeriodOfTime_GetStringValue(t *testing.T) {
	type args struct {
		from   time.Time
		to     time.Duration
		format PeriodOfTimeFormat
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"April 22th, 1997 at 8:54pm + 1d (explicit)", args{
			from:   time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to:     86400 * 1e9,
			format: ExplicitPeriodOfTime}, "19970422T205400Z/19970423T205400Z"},
		{"April 22th, 1997 at 8:54pm + 1d (implicit)", args{
			from:   time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to:     86400 * 1e9,
			format: ImplicitPeriodOfTime}, "19970422T205400Z/P1D"},
		{"April 22th, 1997 at 8:54pm + 1h (implicit)", args{
			from:   time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
			to:     3600 * 1e9,
			format: ImplicitPeriodOfTime}, "19970422T205400Z/PT1H"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			periodOfTime := NewPeriodOfTimeValue(tt.args.from, tt.args.to, tt.args.format)
			if got := periodOfTime.GetStringValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("periodOfTime.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPeriodOfTime_GetValue(t *testing.T) {
	type args struct {
		value string
	}
	type want struct {
		from   time.Time
		to     time.Time
		format PeriodOfTimeFormat
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"April 22th, 1997 at 8:54pm + 1d (explicit)",
			args{value: "19970422T205400Z/19970423T205400Z"},
			want{
				from:   time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
				to:     time.Date(1997, 4, 23, 20, 54, 0, 0, time.UTC),
				format: ExplicitPeriodOfTime}},
		{"April 22th, 1997 at 8:54pm + 1d (implicit)",
			args{value: "19970422T205400Z/P1D"},
			want{
				from:   time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
				to:     time.Date(1997, 4, 23, 20, 54, 0, 0, time.UTC),
				format: ImplicitPeriodOfTime}},
		{"April 22th, 1997 at 8:54pm + 1h (implicit)",
			args{value: "19970422T205400Z/PT1H"}, want{
				from:   time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC),
				to:     time.Date(1997, 4, 22, 21, 54, 0, 0, time.UTC),
				format: ImplicitPeriodOfTime}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			periodOfTime := NewPeriodOfTimeValueFromString(tt.args.value)
			if from, _ := periodOfTime.GetValue(); !reflect.DeepEqual(from, tt.want.from) {
				t.Errorf("(from) periodOfTime.GetValue() = %v, want %v", from, tt.want.from)
			}
			if _, to := periodOfTime.GetValue(); !reflect.DeepEqual(to, tt.want.to) {
				t.Errorf("(to) periodOfTime.GetValue() = %v, want %v", to, tt.want.to)
			}
			if format := periodOfTime.GetFormat(); !reflect.DeepEqual(format, tt.want.format) {
				t.Errorf("(format) periodOfTime.GetFormat() = %v, want %v", format, tt.want.format)
			}
		})
	}
}
