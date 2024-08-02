package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type CommonNameParam interface {
	TextTypeParameter
}

// NewCommonNameParam create a new registries.CommonName property
// This parameter can be used in these properties :
// - registries.Attendee
// - registries.Organizer
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2
func NewCommonNameParam(value string) CommonNameParam {
	return &textParameter{
		ParamName: registries.CommonName,
		Value:     types.NewTextValue(value),
	}
}
