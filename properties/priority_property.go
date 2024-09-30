package properties

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

// NewPriorityProperty create a new registries.PriorityProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.9
func NewPriorityProperty(value int32, params ...parameters.Parameter) PriorityProperty {
	return &integerPropertyType{
		PropName:   registries.PriorityProp,
		Value:      types.NewIntegerValue(value),
		Parameters: params,
	}
}

// NewPriorityPropertyFromString create a new registries.PriorityProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.9
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
