package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type BlockDelimiterProperty interface {
	TextPropertyType
}

func NewBlockDelimiterProperty(
	block registries.PropertyNames,
	component types.BlockDelimiterType,
) BlockDelimiterProperty {
	return &blockDelimiterPropertyType{
		PropName: block,
		Value:    types.NewBlockDelimiterValue(component),
	}
}

func NewBlockDelimiterPropertyFromString(
	block registries.PropertyNames,
	stringComponent string,
) BlockDelimiterProperty {
	return &blockDelimiterPropertyType{
		PropName: block,
		Value:    types.NewBlockDelimiterValue(types.BlockDelimiterType(stringComponent)),
	}
}
