package values

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
)

type DateTimeValue interface {
}
type dateTimeValue struct {
	Value     *time.Time
	IANAToken registries.Type
}

func NewDateTimeValue(value *time.Time) DateTimeValue {
	return &dateTimeValue{
		IANAToken: registries.DATETIME,
		Value:     value,
	}
}
