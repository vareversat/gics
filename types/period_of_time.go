package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9

// Example : 19970101T180000Z/19970102T070000Z (The period starting at 18:00:00 UTC, on January 1, 1997 and ending at 07:00:00 UTC on January 2, 1997)

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/registry"
)

type PeriodValue struct {
	V
	StartValue time.Time
	EndValue   time.Time
	IANAToken  registry.Type
}

func NewPeriodValue(startValue time.Time, endValue time.Time) PeriodValue {
	return PeriodValue{
		IANAToken:  registry.PERIOD,
		StartValue: startValue,
		EndValue:   endValue,
	}
}

func NewPeriodValues(startValues []time.Time, endValues []time.Time) []PeriodValue {
	periodValues := make([]PeriodValue, 0)

	for i := 0; i < len(startValues); i++ {
		periodValues = append(periodValues, PeriodValue{
			V:          NewValue(registry.PERIOD),
			StartValue: startValues[i],
			EndValue:   endValues[i],
		})
	}
	return periodValues
}

func (bV *PeriodValue) GetValue() string {
	return fmt.Sprintf(
		"%s/%s",
		bV.StartValue.Format("20060102T150405Z"),
		bV.EndValue.Format("20060102T150405Z"),
	)
}
