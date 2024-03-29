package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.7

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type LocationProperty interface{}

type locationProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewLocationProperty(value string) LocationProperty {
	return &locationProperty{
		IANAToken: registries.LOCATION,
		Value:     values.NewTextValue(value),
	}
}
