package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.5

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type RelatedToProperty interface {
	TextPropertyType
}

func NewRelatedToProperty(value string, params ...parameters.Parameter) RelatedToProperty {
	return &textPropertyType{
		PropName:   registries.RELATEDTO,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
