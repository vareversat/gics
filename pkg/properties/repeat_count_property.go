package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.2

import (
	"fmt"
	"strconv"

	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type RepeatCountProperty interface {
	IntegerPropertyType
}

func NewRepeatCountProperty(value int32, params ...parameters.Parameter) RepeatCountProperty {
	return &integerPropertyType{
		PropName:   registries.REPEAT,
		Value:      types.NewIntegerValue(value),
		Parameters: params,
	}
}

func NewRepeatCountPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (RepeatCountProperty, error) {
	repeat, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s cannot be parsed as int32", value)
	}
	return &integerPropertyType{
		PropName:   registries.REPEAT,
		Value:      types.NewIntegerValue(int32(repeat)),
		Parameters: params,
	}, nil
}
