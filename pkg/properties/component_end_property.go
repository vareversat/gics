package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type ComponentEndProperty interface {
}

type componentEndProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewComponentEndProperty(component registries.Component) ComponentEndProperty {
	return &componentEndProperty{
		IANAToken: registries.END,
		Value:     values.NewTextValue(string(component)),
	}
}
