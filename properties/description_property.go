package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type DescriptionProperty interface {
	TextPropertyType
}

func NewDescriptionProperty(descriptionValue string,
	params ...parameters.Parameter) DescriptionProperty {
	return &textPropertyType{
		PropName:   registry.DESCRIPTION,
		Value:      types.NewTextValue(descriptionValue),
		Parameters: params,
	}
}
