package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.5

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RelatedToProperty interface {
	TextPropertyType
}

func NewRelatedToProperty(value string, params ...parameters.Parameter) RelatedToProperty {
	return &textPropertyType{
		PropName:   registries.RelatedTo,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
