package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type DescriptionProperty interface{}

type descriptionProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewDescriptionProperty(value string) DescriptionProperty {
	return &descriptionProperty{
		IANAToken: registries.DESCRIPTION,
		Value:     values.NewTextValue(value),
	}
}
