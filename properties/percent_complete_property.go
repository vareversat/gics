package properties

import (
	"fmt"
	"strconv"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type PercentCompleteProperty interface {
	IntegerPropertyType
}

// NewPercentCompleteProperty create a new registries.PercentCompleteProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtodo (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.8
func NewPercentCompleteProperty(value int32, params ...parameters.Parameter) IntegerPropertyType {
	return &integerPropertyType{
		PropName:   registries.PercentCompleteProp,
		Value:      types.NewIntegerValue(value),
		Parameters: params,
	}
}

// NewPercentCompletePropertyFromString create a new registries.PercentCompleteProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtodo (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.8
func NewPercentCompletePropertyFromString(
	value string,
	params ...parameters.Parameter,
) (IntegerPropertyType, error) {
	percentage, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s cannot be parsed as int32", value)
	}
	return &integerPropertyType{
		PropName:   registries.PercentCompleteProp,
		Value:      types.NewIntegerValue(int32(percentage)),
		Parameters: params,
	}, nil
}
