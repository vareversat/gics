package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.1

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type CalScaleProperty interface{}

type calScaleProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewCalScaleProperty() CalScaleProperty {
	return &calScaleProperty{
		IANAToken: registries.CALSCALE,
		Value:     "GREGORIAN", // This is the default value of the CALSCALE property
	}
}
