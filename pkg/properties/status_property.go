package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type StatusProperty interface {
}

type statusProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewStatusProperty(status registries.Status) StatusProperty {
	return &statusProperty{
		IANAToken: registries.STATUS,
		Value:     values.NewTextValue(string(status)),
	}
}
