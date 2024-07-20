package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type ClassificationProperty interface {
	TextPropertyType
}

func NewClassificationProperty(
	classValue registry.ClassificationRegistry,
	params ...parameters.Parameter,
) ClassificationProperty {
	return &classificationPropertyType{
		PropName:   registry.CLASS,
		Value:      types.NewClassificationValue(classValue),
		Parameters: params,
	}
}

func NewClassificationPropertyFromString(
	classValue string,
	params ...parameters.Parameter,
) ClassificationProperty {
	return &classificationPropertyType{
		PropName:   registry.CLASS,
		Value:      types.NewClassificationValue(registry.ClassificationRegistry(classValue)),
		Parameters: params,
	}
}
