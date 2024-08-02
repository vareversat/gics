package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type FormatTypeParam interface {
	TextTypeParameter
}

// NewFormatTypeParam create a new registries.FormatType property
// This parameter can be used in this property :
// - registries.Attachment
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.8
func NewFormatTypeParam(value string) FormatTypeParam {
	return &textParameter{
		ParamName: registries.FormatType,
		Value:     types.NewTextValue(value),
	}
}
