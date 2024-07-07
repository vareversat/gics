package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ParticipationRoleParam interface {
	TextParameter
}

func NewParticipationRoleParam(value ParticipationRoleType) ParticipationRoleParam {
	return &textParameter{
		ParamName: registries.ROLE,
		Value:     types.NewTextValue(string(value)),
	}
}

type ParticipationRoleType string

const (
	CHAIR          ParticipationRoleType = "CHAIR"
	REQPARTICIPANT ParticipationRoleType = "REQ-PARTICIPANT"
	OPTPARTICIPANT ParticipationRoleType = "OPT-PARTICIPANT"
	NONPARTICIPANT ParticipationRoleType = "NON-PARTICIPANT"
)
