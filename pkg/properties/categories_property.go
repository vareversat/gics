package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.2

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
	"github.com/vareversat/gics/pkg/utils"
)

type CategoriesProperty interface {
	TextPropertyType
}

func NewCategoryProperty(values []string, params ...parameters.Parameter) CategoriesProperty {
	return &textPropertyType{
		PropName:   registries.CATEGORIES,
		Values:     types.NewTextValues(values),
		Parameters: params,
	}
}

func NewCategoryPropertyFromString(
	values string,
	params ...parameters.Parameter,
) CategoriesProperty {
	return &textPropertyType{
		PropName:   registries.CATEGORIES,
		Values:     types.NewTextValues(utils.StringToStringArray(values)),
		Parameters: params,
	}
}
