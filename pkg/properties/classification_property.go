package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type ClassificationProperty interface {
}

type classificationProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewClassificationProperty(classification registries.Classification) ClassificationProperty {
	return &classificationProperty{
		IANAToken: registries.CLASS,
		Value:     values.NewTextValue(string(classification)),
	}
}
