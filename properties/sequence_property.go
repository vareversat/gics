package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4

import (
	"fmt"
	"strconv"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type SequenceProperty interface {
	IntegerPropertyType
}

func NewSequenceProperty(value int32, params ...parameters.Parameter) SequenceProperty {
	return &integerPropertyType{
		PropName:   registries.SEQUENCE,
		Value:      types.NewIntegerValue(value),
		Parameters: params,
	}
}

func NewSequencePropertyFromString(
	value string,
	params ...parameters.Parameter,
) (SequenceProperty, error) {
	sequence, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s cannot be parsed as int32", value)
	}
	return &integerPropertyType{
		PropName:   registries.SEQUENCE,
		Value:      types.NewIntegerValue(int32(sequence)),
		Parameters: params,
	}, nil
}
