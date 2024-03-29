package values

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
)

type DateValue interface{}

type dateValue struct {
	Value     *time.Time
	IANAToken registries.Type
}

func NewDateValue(value *time.Time) DateValue {
	return &dateValue{
		IANAToken: registries.DATE,
		Value:     value,
	}
}
