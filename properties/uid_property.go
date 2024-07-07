package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type UidProperty interface {
	TextPropertyType
}

func NewUidProperty(value string, params ...parameters.Parameter) UidProperty {
	return &textPropertyType{
		PropName:   registries.UID,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
