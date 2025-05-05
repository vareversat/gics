package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewBySecondPart(t *testing.T) {
	type args struct {
		seconds []int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Create BYSECOND Recurrence rule component #1", args{seconds: []int32{10}}, "10", false},
		{
			"Create BYSECOND Recurrence rule component #2",
			args{seconds: []int32{10, 11}},
			"10,11",
			false,
		},
		{
			"Create BYSECOND Recurrence rule component #3",
			args{seconds: []int32{10, 22, 100}},
			"0",
			true,
		},
		{
			"Create BYSECOND Recurrence rule component #4",
			args{seconds: []int32{10, 22, -100}},
			"0",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBySecondPart(tt.args.seconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBySecondPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}

func TestNewBySecondPartFromString(t *testing.T) {
	want, _ := NewBySecondPart([]int32{10, 11})
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    BySecondPart
		wantErr bool
	}{
		{"Create BySecond (no error)", args{value: "10, 11"}, want, false},
		{"Create BySecond (with error)", args{value: "100"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBySecondPartFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("got \"%s\" error, want not", err.Error())
			} else if got != nil && !reflect.DeepEqual(got.GetPartValue(), tt.want.GetPartValue()) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want.GetPartValue())
			}
		})
	}
}
