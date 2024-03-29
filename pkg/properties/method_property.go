package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.2

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type MethodProperty interface {
}

type methodProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewMethodProperty(value string) MethodProperty {
	return &methodProperty{
		IANAToken: registries.METHOD,
		Value:     values.NewTextValue(value),
	}
}
