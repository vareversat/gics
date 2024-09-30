package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.2

import (
	"fmt"
	"strconv"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RepeatCountProperty interface {
	IntegerPropertyType
}

// NewRepeatCountProperty create a new registries.RepeatProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.2
func NewRepeatCountProperty(value int32, params ...parameters.Parameter) RepeatCountProperty {
	return &integerPropertyType{
		PropName:   registries.RepeatProp,
		Value:      types.NewIntegerValue(value),
		Parameters: params,
	}
}

// NewRepeatCountPropertyFromString create a new registries.RepeatProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.2
func NewRepeatCountPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (RepeatCountProperty, error) {
	repeat, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s cannot be parsed as int32", value)
	}
	return &integerPropertyType{
		PropName:   registries.RepeatProp,
		Value:      types.NewIntegerValue(int32(repeat)),
		Parameters: params,
	}, nil
}
