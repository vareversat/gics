package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9

// Parameter used in these properties :
// - FREEBUSY

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type FreeBusyTimeParam interface {
	TextTypeParameter
}

func NewFreeBusyTimeParam(value registries.FreeBusyTimeTypeRegistry) FreeBusyTimeParam {
	return &textParameter{
		ParamName: registries.FreeBusyTimeType,
		Value:     types.NewTextValue(string(value)),
	}
}
