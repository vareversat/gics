package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ParticipationRoleParam interface {
	TextTypeParameter
}

// NewParticipationRoleParam create a new registries.RoleParam property
// This parameter can be used in this property :
// - registries.AttendeeProp
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
func NewParticipationRoleParam(value registries.ParticipantRoleRegistry) ParticipationRoleParam {
	return &textParameter{
		ParamName: registries.RoleParam,
		Value:     types.NewTextValue(string(value)),
	}
}
