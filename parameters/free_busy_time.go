package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type FreeBusyTimeParam interface {
	TextTypeParameter
}

// NewFreeBusyTimeParam create a new registries.FreeBusyTimeTypeParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.FreeBusyTimeProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9
func NewFreeBusyTimeParam(value registries.FreeBusyTimeTypeRegistry) FreeBusyTimeParam {
	return &textParameter{
		ParamName: registries.FreeBusyTimeTypeParam,
		Value:     types.NewTextValue(string(value)),
	}
}
