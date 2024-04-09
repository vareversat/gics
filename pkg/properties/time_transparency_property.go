package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type TimeTransparencyProperty interface {
	TimeTransparencyPropertyType
}

func NewTimeTransparencyProperty(
	value types.TimeTransparencyType,
	params ...parameters.Parameter,
) TimeTransparencyProperty {
	return &timeTransparencyPropertyType{
		PropName:   registries.TRANSP,
		Value:      types.NewTimeTransparencyValue(value),
		Parameters: params,
	}
}

func NewTimeTransparencyPropertyFromString(
	value string,
	params ...parameters.Parameter,
) TimeTransparencyProperty {
	return &timeTransparencyPropertyType{
		PropName:   registries.TRANSP,
		Value:      types.NewTimeTransparencyValue(types.TimeTransparencyType(value)),
		Parameters: params,
	}
}
