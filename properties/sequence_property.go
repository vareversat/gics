package properties

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

// NewSequenceProperty create a new registries.SequenceProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4
func NewSequenceProperty(value int32, params ...parameters.Parameter) SequenceProperty {
	return &integerPropertyType{
		PropName:   registries.SequenceProp,
		Value:      types.NewIntegerValue(value),
		Parameters: params,
	}
}

// NewSequencePropertyFromString create a new registries.SequenceProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4
func NewSequencePropertyFromString(
	value string,
	params ...parameters.Parameter,
) (SequenceProperty, error) {
	sequence, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s cannot be parsed as int32", value)
	}
	return &integerPropertyType{
		PropName:   registries.SequenceProp,
		Value:      types.NewIntegerValue(int32(sequence)),
		Parameters: params,
	}, nil
}
