package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type RSVPParam interface {
	BooleanTypeParameter
}

func NewRSVPParam(value bool) RSVPParam {
	return &booleanParameter{
		ParamName: registry.RSVP,
		Value:     types.NewBooleanValue(value),
	}
}
