package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewByMonthPart(t *testing.T) {
	type args struct {
		months []int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Create BYMONTH Recurrence rule component #1", args{months: []int32{10}}, "10", false},
		{
			"Create BYMONTH Recurrence rule component #2",
			args{months: []int32{10, 11}},
			"10,11",
			false,
		},
		{
			"Create BYMONTH Recurrence rule component #3",
			args{months: []int32{10, 22, 100}},
			"0",
			true,
		},
		{
			"Create BYMONTH Recurrence rule component #4",
			args{months: []int32{10, 22, -100}},
			"0",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByMonthPart(tt.args.months)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByMonthPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}
