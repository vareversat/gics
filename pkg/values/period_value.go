package values

import (
	"github.com/vareversat/gics/pkg/registries"
)

type PeriodValue interface{}

type periodValue struct {
	StartValue DateTimeValue
	EndValue   DateTimeValue
	IANAToken  registries.Type
}

func NewPeriodValue(startValue DateTimeValue, endValue DateTimeValue) PeriodValue {
	return &periodValue{
		IANAToken:  registries.PERIOD,
		StartValue: startValue,
		EndValue:   endValue,
	}
}
