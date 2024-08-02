package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ClassificationProperty interface {
	TextPropertyType
}

func NewClassificationProperty(
	classValue registries.ClassificationRegistry,
	params ...parameters.Parameter,
) ClassificationProperty {
	return &classificationPropertyType{
		PropName:   registries.ClassProp,
		Value:      types.NewClassificationValue(classValue),
		Parameters: params,
	}
}

func NewClassificationPropertyFromString(
	classValue string,
	params ...parameters.Parameter,
) ClassificationProperty {
	return &classificationPropertyType{
		PropName:   registries.ClassProp,
		Value:      types.NewClassificationValue(registries.ClassificationRegistry(classValue)),
		Parameters: params,
	}
}
