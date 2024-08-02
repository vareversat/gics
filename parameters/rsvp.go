package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RSVPParam interface {
	BooleanTypeParameter
}

// NewRSVPParam create a new registries.RsvpExpectation property
// This parameter can be used in this property :
// - registries.Attendee
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17
func NewRSVPParam(value bool) RSVPParam {
	return &booleanParameter{
		ParamName: registries.RsvpExpectation,
		Value:     types.NewBooleanValue(value),
	}
}
