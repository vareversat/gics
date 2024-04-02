package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type DescriptionProperty interface {
	TextPropertyType
}

func NewDescriptionProperty(descriptionValue string,
	params ...parameters.Parameter) DescriptionProperty {
	return &textPropertyType{
		PropName:   registries.DESCRIPTION,
		Value:      types.NewTextValue(descriptionValue),
		Parameters: params,
	}
}
