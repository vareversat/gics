package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type CommonNameParam interface {
	TextTypeParameter
}

// NewCommonNameParam create a new registries.CommonNameParam property. See [RFC-5545] ref for more info
// This parameter can be used in these properties :
// - registries.AttendeeProp
// - registries.OrganizerProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2
func NewCommonNameParam(value string) CommonNameParam {
	return &textParameter{
		ParamName: registries.CommonNameParam,
		Value:     types.NewTextValue(value),
	}
}
