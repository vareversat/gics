package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.2

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type CategoriesProperty interface{}

type categoriesProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewCategoryProperty(value string) CategoriesProperty {
	return &categoriesProperty{
		IANAToken: registries.CATEGORIES,
		Value:     values.NewTextValue(value),
	}
}
