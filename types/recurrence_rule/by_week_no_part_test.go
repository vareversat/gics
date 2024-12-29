package recurrence_rule

import (
	"reflect"
	"testing"
)

func TestNewByWeekNoPart(t *testing.T) {
	type args struct {
		weekNumbers []int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Create BYWEEKNO Recurrence rule component #1",
			args{weekNumbers: []int32{10}},
			"10",
			false,
		},
		{
			"Create BYWEEKNO Recurrence rule component #2",
			args{weekNumbers: []int32{10, 11}},
			"10,11",
			false,
		},
		{
			"Create BYWEEKNO Recurrence rule component #3",
			args{weekNumbers: []int32{10, 22, 100}},
			"0",
			true,
		},
		{
			"Create BYWEEKNO Recurrence rule component #4",
			args{weekNumbers: []int32{10, 22, -100}},
			"0",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByWeekNoPart(tt.args.weekNumbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByWeekNoPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetPartValue(), tt.want) {
				t.Errorf("got.ToString() = %s, want %s", got.GetPartValue(), tt.want)
			}
		})
	}
}
