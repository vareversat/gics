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
