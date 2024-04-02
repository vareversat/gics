package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type ProductIdProperty interface {
	TextPropertyType
}

func NewProductIdProperty(value string) ProductIdProperty {
	return &textPropertyType{
		PropName: registries.PRODID,
		Value:    types.NewTextValue(value),
	}
}
