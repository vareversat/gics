package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.2

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
	"github.com/vareversat/gics/utils"
)

type CategoriesProperty interface {
	TextPropertyType
}

func NewCategoryProperty(values []string, params ...parameters.Parameter) CategoriesProperty {
	return &textPropertyType{
		PropName:   registries.CategoriesProp,
		Values:     types.NewTextValues(values),
		Parameters: params,
	}
}

func NewCategoryPropertyFromString(
	values string,
	params ...parameters.Parameter,
) CategoriesProperty {
	return &textPropertyType{
		PropName:   registries.CategoriesProp,
		Values:     types.NewTextValues(utils.StringToStringArray(values)),
		Parameters: params,
	}
}
