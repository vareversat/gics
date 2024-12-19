package types

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTimeValue(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Paris")
	type args struct {
		value time.Time
	}
	tests := []struct {
		name string
		args args
		want TimeFormat
	}{
		{"Time local time", args{value: time.Now().In(loc)}, TimeWithLocalTime},
		{"Time UTC", args{value: time.Now()}, TimeWithUtcTime},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeValue := NewTimeValue(tt.args.value)
			if got := timeValue.GetFormat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeValue.GetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeTimeGetStringValue(t *testing.T) {
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
		{"Time time local time", args{value: currentTime1}, "225458"},
		{"Time time UTC", args{value: currentTime2}, "225458Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeValue := NewTimeValue(tt.args.value)
			if got := timeValue.GetStringValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeValue.GetStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTimeTimeValueFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want TimeFormat
	}{
		{"Time time local time", args{value: "020000"}, TimeWithLocalTime},
		{"Time time UTC", args{value: "070000Z"}, TimeWithUtcTime},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeValue := NewTimeValueFromString(tt.args.value)
			if got := timeValue.GetFormat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeValue.GetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
