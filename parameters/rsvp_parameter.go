package parameters

import (
	"fmt"
	"strconv"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RSVPParam interface {
	BooleanTypeParameter
}

// NewRSVPParam create a new registries.RsvpExpectationParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.AttendeeProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17
func NewRSVPParam(value bool) RSVPParam {
	return &booleanParameter{
		ParamName: registries.RsvpExpectationParam,
		Value:     types.NewBooleanValue(value),
	}
}

func NewRSVPParamFromString(value string) (RSVPParam, error) {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid boolean value", value)
	}

	return &booleanParameter{
		ParamName: registries.RsvpExpectationParam,
		Value:     types.NewBooleanValue(boolValue),
	}, nil
}
