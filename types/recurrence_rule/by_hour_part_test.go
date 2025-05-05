package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewByHourPart(t *testing.T) {
	type args struct {
		hours []int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Create BYHOUR Recurrence rule component #1", args{hours: []int32{10}}, "10", false},
		{
			"Create BYHOUR Recurrence rule component #2",
			args{hours: []int32{10, 22}},
			"10,22",
			false,
		},
		{"Create BYHOUR Recurrence rule component #3", args{hours: []int32{10, 22, 33}}, "0", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByHourPart(tt.args.hours)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByHourPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}

func TestNewByHourPartFromString(t *testing.T) {
	want, _ := NewByHourPart([]int32{18, 19})
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    ByHourPart
		wantErr bool
	}{
		{"Create ByHour (no error)", args{value: "18,  19  "}, want, false},
		{"Create ByHour (with error)", args{value: "100"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByHourPartFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("got \"%s\" error, want not", err.Error())
			} else if got != nil && !reflect.DeepEqual(got.GetPartValue(), tt.want.GetPartValue()) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want.GetPartValue())
			}
		})
	}
}
