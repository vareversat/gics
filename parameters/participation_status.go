package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ParticipationStatusParam interface {
	TextTypeParameter
}

// NewParticipationStatusParam create a new registries.ParticipationStatusParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.AttendeeProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
func NewParticipationStatusParam(
	value registries.ParticipantStatusRegistry,
) ParticipationStatusParam {
	return &textParameter{
		ParamName: registries.ParticipationStatusParam,
		Value:     types.NewTextValue(string(value)),
	}
}
