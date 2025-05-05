package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewBySetPosPart(t *testing.T) {
	type args struct {
		positions []int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Create BYSETPOS Recurrence rule component #1", args{positions: []int32{10}}, "10", false},
		{
			"Create BYSETPOS Recurrence rule component #2",
			args{positions: []int32{10, 11}},
			"10,11",
			false,
		},
		{
			"Create BYSETPOS Recurrence rule component #3",
			args{positions: []int32{10, 22, 400}},
			"0",
			true,
		},
		{
			"Create BYSETPOS Recurrence rule component #4",
			args{positions: []int32{10, 22, -400}},
			"0",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBySetPosPart(tt.args.positions)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBySetPosPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}

func TestNewBySetPosPartFromString(t *testing.T) {
	want, _ := NewBySetPosPart([]int32{10, 11, 200})
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    BySetPosPart
		wantErr bool
	}{
		{"Create BySetPos (no error)", args{value: "10, 11, 200"}, want, false},
		{"Create BySetPos (with error)", args{value: "1000"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBySetPosPartFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("got \"%s\" error, want not", err.Error())
			} else if got != nil && !reflect.DeepEqual(got.GetPartValue(), tt.want.GetPartValue()) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want.GetPartValue())
			}
		})
	}
}
