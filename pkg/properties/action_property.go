package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type ActionProperty interface{}

type actionProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewActionProperty(action registries.Action) ActionProperty {
	return &actionProperty{
		IANAToken: registries.ACTION,
		Value:     values.NewTextValue(string(action)),
	}
}
