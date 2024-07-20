package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9

// Parameter used in these properties :
// - FREEBUSY

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type FreeBusyTimeParam interface {
	TextTypeParameter
}

func NewFreeBusyTimeParam(value registry.FreeBusyTimeTypeRegistry) FreeBusyTimeParam {
	return &textParameter{
		ParamName: registry.FreeBusyTimeType,
		Value:     types.NewTextValue(string(value)),
	}
}
