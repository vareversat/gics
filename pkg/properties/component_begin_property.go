package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type ComponentBeginProperty interface {
}

type componentBeginProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewComponentBeginProperty(component registries.Component) ComponentBeginProperty {
	return &componentBeginProperty{
		IANAToken: registries.BEGIN,
		Value:     values.NewTextValue(string(component)),
	}
}
