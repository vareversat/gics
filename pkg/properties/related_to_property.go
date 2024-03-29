package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.5

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type RelatedToProperty interface {
}

type relatedToProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewRelatedToProperty(value string) RelatedToProperty {
	return &relatedToProperty{
		IANAToken: registries.RELATEDTO,
		Value:     values.NewTextValue(value),
	}
}
