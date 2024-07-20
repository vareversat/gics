package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2

// Parameter used in these properties :
// - ATTENDEE
// - ORGANIZER

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type CommonNameParam interface {
	TextTypeParameter
}

type commonNameParam struct {
	IANAToken registry.ParameterRegistry
	Value     types.TextValue
}

func NewCommonNameParam(value string) CommonNameParam {
	return &textParameter{
		ParamName: registry.CommonName,
		Value:     types.NewTextValue(value),
	}
}
