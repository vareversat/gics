package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.9

import (
	"fmt"
	"strconv"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type PriorityProperty interface {
	IntegerPropertyType
}

func NewPriorityProperty(value int32, params ...parameters.Parameter) PriorityProperty {
	return &integerPropertyType{
		PropName:   registries.PriorityProp,
		Value:      types.NewIntegerValue(value),
		Parameters: params,
	}
}

func NewPriorityPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (PriorityProperty, error) {
	priority, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s cannot be parsed as int32", value)
	}
	return &integerPropertyType{
		PropName:   registries.PriorityProp,
		Value:      types.NewIntegerValue(int32(priority)),
		Parameters: params,
	}, nil
}
