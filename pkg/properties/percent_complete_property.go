package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.8

import (
	"fmt"
	"strconv"

	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type PercentCompleteProperty interface {
	IntegerPropertyType
}

func NewPercentCompleteProperty(value int32, params ...parameters.Parameter) IntegerPropertyType {
	return &integerPropertyType{
		PropName:   registries.PERCENTCOMPLETE,
		Value:      types.NewIntegerValue(value),
		Parameters: params,
	}
}

func NewPercentCompletePropertyFromString(
	value string,
	params ...parameters.Parameter,
) (IntegerPropertyType, error) {
	percentage, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s cannot be parsed as int32", value)
	}
	return &integerPropertyType{
		PropName:   registries.PERCENTCOMPLETE,
		Value:      types.NewIntegerValue(int32(percentage)),
		Parameters: params,
	}, nil
}
