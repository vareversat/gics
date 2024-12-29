package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewByYearDayPart(t *testing.T) {
	type args struct {
		days []int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Create BYYEARDAY Recurrence rule component #1", args{days: []int32{10}}, "10", false},
		{
			"Create BYYEARDAY Recurrence rule component #2",
			args{days: []int32{10, 11}},
			"10,11",
			false,
		},
		{
			"Create BYYEARDAY Recurrence rule component #3",
			args{days: []int32{10, 22, 100}},
			"10,22,100",
			false,
		},
		{
			"Create BYYEARDAY Recurrence rule component #4",
			args{days: []int32{10, 22, -400}},
			"0",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByYearDayPart(tt.args.days)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByYearDayPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}
