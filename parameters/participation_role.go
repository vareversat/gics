package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type ParticipationRoleParam interface {
	TextTypeParameter
}

func NewParticipationRoleParam(value registry.ParticipantRoleRegistry) ParticipationRoleParam {
	return &textParameter{
		ParamName: registry.Role,
		Value:     types.NewTextValue(string(value)),
	}
}
