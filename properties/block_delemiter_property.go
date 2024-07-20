package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type BlockDelimiterProperty interface {
	TextPropertyType
}

func NewBlockDelimiterProperty(
	block registry.PropertyNames,
	component types.BlockDelimiterType,
) BlockDelimiterProperty {
	return &blockDelimiterPropertyType{
		PropName: block,
		Value:    types.NewBlockDelimiterValue(component),
	}
}

func NewBlockDelimiterPropertyFromString(
	block registry.PropertyNames,
	stringComponent string,
) BlockDelimiterProperty {
	return &blockDelimiterPropertyType{
		PropName: block,
		Value:    types.NewBlockDelimiterValue(types.BlockDelimiterType(stringComponent)),
	}
}
