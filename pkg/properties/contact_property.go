package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type ContactProperty interface {
	TextPropertyType
}

func NewContactProperty(value string, params ...parameters.Parameter) ContactProperty {
	return &textPropertyType{
		PropName:   registries.CONTACT,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
