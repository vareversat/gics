package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RSVPParam interface {
	BooleanParameter
}

func NewRSVPParam(value bool) RSVPParam {
	return &booleanParameter{
		ParamName: registries.RSVP,
		Value:     types.NewBooleanValue(value),
	}
}
