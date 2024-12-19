package types

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDateTimeValue(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Paris")
	type args struct {
		value time.Time
	}
	tests := []struct {
		name string
		args args
		want DateTimeFormat
	}{
		{"Date time local time", args{value: time.Now().In(loc)}, WithLocalTime},
		{"Date time UTC", args{value: time.Now()}, WithUtcTime},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataTime := NewDateTimeValue(tt.args.value)
			if got := dataTime.GetFormat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataTime.GetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTimeGetStringValue(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Paris")
	currentTime1 := time.Date(2024, 10, 15, 22, 54, 58, 651387237, loc)
	currentTime2 := time.Date(2024, 10, 15, 22, 54, 58, 651387237, time.UTC)

	type args struct {
		value time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Date time local time", args{value: currentTime1}, "20241015T225458"},
		{"Date time UTC", args{value: currentTime2}, "20241015T225458Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataTime := NewDateTimeValue(tt.args.value)
			if got := dataTime.GetStringValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataTime.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDateTimeValueFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want DateTimeFormat
	}{
		{"Date time local time", args{value: "19980119T020000"}, WithLocalTime},
		{"Date time UTC", args{value: "19980119T070000Z"}, WithUtcTime},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataTime := NewDateTimeValueFromString(tt.args.value)
			if got := dataTime.GetFormat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataTime.GetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
