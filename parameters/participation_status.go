package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ParticipationStatusParam interface {
	TextTypeParameter
}

func NewParticipationStatusParam(
	value registries.ParticipantStatusRegistry,
) ParticipationStatusParam {
	return &textParameter{
		ParamName: registries.ParticipationStatus,
		Value:     types.NewTextValue(string(value)),
	}
}
