package properties

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

func TestNewRecurrenceDateTimesProperty_NoParam(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name     string
		args     args
		errorMsg string
	}{
		{"Create new RDATE from DATE (no VALUE param)", args{
			date: time.Date(
				1997,
				4,
				22,
				20,
				54,
				0,
				0,
				time.UTC,
			)}, "you have to explicitly set the value type via the VALUE parameter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRecurrenceDateTimesFromDateProperty(tt.args.date)
			assert.EqualError(
				t,
				err,
				"you have to explicitly set the value type via the VALUE parameter",
			)
		})
	}
}

func TestNewRecurrenceDateTimesProperty_WrongParam(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name     string
		args     args
		errorMsg string
	}{
		{"Create new RDATE from DATE (wrong VALUE param)", args{
			date: time.Date(
				1997,
				4,
				22,
				20,
				54,
				0,
				0,
				time.UTC,
			)}, "you have to explicitly set the value type via the VALUE parameter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRecurrenceDateTimesFromDateProperty(
				tt.args.date,
				parameters.NewValueDataTypesParam(registries.DateTime),
			)
			assert.EqualError(t, err, "VALUE parameter must be set to 'DATE'")
		})
	}
}

func TestNewRecurrenceDateTimesProperty_Date(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Create new RDATE from DATE", args{
			date: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC)}, "19970422"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rDate, _ := NewRecurrenceDateTimesFromDateProperty(
				tt.args.date,
				parameters.NewValueDataTypesParam(registries.Date),
			)
			if got := rDate.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("periodOfTime.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRecurrenceDateTimesProperty_DateTime(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Create new RDATE from DATE-TIME", args{
			date: time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC)}, "19970422T205400Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rDate, _ := NewRecurrenceDateTimesFromDateTimeProperty(tt.args.date)
			if got := rDate.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("periodOfTime.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRecurrenceDateTimesProperty_ImplicitPeriod(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Create new RDATE from implicit PERIOD", args{
			from: time.Date(
				1997,
				4,
				22,
				20,
				54,
				0,
				0,
				time.UTC,
			), to: 3600 * 1e9}, "19970422T205400Z/PT1H"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rDate, _ := NewRecurrenceDateTimesFromImplicitPeriodProperty(
				tt.args.from,
				tt.args.to,
				parameters.NewValueDataTypesParam(registries.PeriodOfTime),
			)
			if got := rDate.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("periodOfTime.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestNewRecurrenceDateTimesPropertyFromString(t *testing.T) {
	type args struct {
		value string
		param registries.ValueTypeRegistry
	}
	tests := []struct {
		name string
		args args
		want types.ValueType
	}{
		{"Create new RDATE from implicit PERIOD", args{
			value: "19970422T205400Z/PT1H", param: registries.PeriodOfTime}, types.NewImplicitPeriodOfTimeValue(time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC), 3600*1e9)},
		{"Create new RDATE from explicit PERIOD", args{
			value: "19970422T205400Z/19970423T205400Z", param: registries.PeriodOfTime}, types.NewExplicitPeriodOfTimeValue(time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC), time.Date(1997, 4, 23, 20, 54, 0, 0, time.UTC))},
		{"Create new RDATE from DATE", args{
			value: "19970422", param: registries.Date}, types.NewDateValue(time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC))},
		{"Create new RDATE from DATE-TIME", args{
			value: "19970422T205400Z", param: registries.DateTime}, types.NewDateTimeValue(time.Date(1997, 4, 22, 20, 54, 0, 0, time.UTC))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rDate, _ := NewRecurrenceDateTimesPropertyFromString(
				tt.args.value,
				parameters.NewValueDataTypesParam(tt.args.param),
			)
			if got := rDate.GetValue(); !reflect.DeepEqual(got, tt.want.GetStringValue()) {
				t.Errorf("rDate.GetValue() = %v, want %v", got, tt.want.GetStringValue())
			}
		})
	}
}
