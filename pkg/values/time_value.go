package values

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
)

type TimeValue interface {
}
type timeValue struct {
	Value     *time.Time
	IANAToken registries.Type
}

func NewTimeValue(value *time.Time) TimeValue {
	return &timeValue{
		IANAToken: registries.TIME,
		Value:     value,
	}
}
