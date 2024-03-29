package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type ProIdProperty interface {
}

type proIdProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewProIdProperty(value string) ProIdProperty {
	return &proIdProperty{
		IANAToken: registries.PROID,
		Value:     values.NewTextValue(value),
	}
}
