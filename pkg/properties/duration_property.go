package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type DurationProperty interface {
}

type durationProperty struct {
	IANAToken registries.Properties
	Value     values.PeriodValue
}

func NewPeriodProperty(startValue *time.Time, endValue *time.Time) DurationProperty {
	return &durationProperty{
		IANAToken: registries.DURATION_PROP,
		Value:     values.NewPeriodValue(startValue, endValue),
	}
}
