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
	classValue types.ClassificationType,
	params ...parameters.Parameter,
) ClassificationProperty {
	return &classificationPropertyType{
		PropName:   registries.CLASS,
		Value:      types.NewClassificationValue(classValue),
		Parameters: params,
	}
}

func NewClassificationPropertyFromString(
	classValue string,
	params ...parameters.Parameter,
) ClassificationProperty {
	return &classificationPropertyType{
		PropName:   registries.CLASS,
		Value:      types.NewClassificationValue(types.ClassificationType(classValue)),
		Parameters: params,
	}
}
