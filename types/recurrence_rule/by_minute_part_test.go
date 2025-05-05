package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewByMinutePart(t *testing.T) {
	type args struct {
		minutes []int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Create BYMINUTE Recurrence rule component #1", args{minutes: []int32{10}}, "10", false},
		{
			"Create BYMINUTE Recurrence rule component #2",
			args{minutes: []int32{10, 22}},
			"10,22",
			false,
		},
		{
			"Create BYMINUTE Recurrence rule component #3",
			args{minutes: []int32{10, 22, 100}},
			"0",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByMinutePart(tt.args.minutes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByMinutePart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}

func TestNewByMinutePartFromString(t *testing.T) {
	want, _ := NewByMinutePart([]int32{18, 19})
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    ByMinutePart
		wantErr bool
	}{
		{"Create ByMinute (no error)", args{value: "18, 19"}, want, false},
		{"Create ByMinute (with error)", args{value: "100"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByMinutePartFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("got \"%s\" error, want not", err.Error())
			} else if got != nil && !reflect.DeepEqual(got.GetPartValue(), tt.want.GetPartValue()) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want.GetPartValue())
			}
		})
	}
}
