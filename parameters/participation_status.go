package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ParticipationStatusParam interface {
	TextTypeParameter
}

// NewParticipationStatusParam create a new registries.ParticipationStatus property
// This parameter can be used in this property :
// - registries.Attendee
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
func NewParticipationStatusParam(
	value registries.ParticipantStatusRegistry,
) ParticipationStatusParam {
	return &textParameter{
		ParamName: registries.ParticipationStatus,
		Value:     types.NewTextValue(string(value)),
	}
}
