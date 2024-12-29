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
