package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ParticipationRoleParam interface {
	TextTypeParameter
}

func NewParticipationRoleParam(value registries.ParticipantRoleRegistry) ParticipationRoleParam {
	return &textParameter{
		ParamName: registries.Role,
		Value:     types.NewTextValue(string(value)),
	}
}
