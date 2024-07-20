package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type ContactProperty interface {
	TextPropertyType
}

func NewContactProperty(value string, params ...parameters.Parameter) ContactProperty {
	return &textPropertyType{
		PropName:   registry.CONTACT,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
