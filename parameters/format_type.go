package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type FormatTypeParam interface {
	TextTypeParameter
}

// NewFormatTypeParam create a new registries.FormatTypeParam property
// This parameter can be used in this property :
// - registries.AttachmentProp
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.8
func NewFormatTypeParam(value string) FormatTypeParam {
	return &textParameter{
		ParamName: registries.FormatTypeParam,
		Value:     types.NewTextValue(value),
	}
}
