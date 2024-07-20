package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.3

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type ProductIdProperty interface {
	TextPropertyType
}

func NewProductIdProperty(value string, params ...parameters.Parameter) ProductIdProperty {
	return &textPropertyType{
		PropName:   registry.PRODID,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
