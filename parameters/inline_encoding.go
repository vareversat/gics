package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type InlineEncodingParam interface {
	TextTypeParameter
}

// NewInlineEncodingParam create a new registries.Encoding property
// This parameter can be used in this property :
// - registries.Binary
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.7
func NewInlineEncodingParam(value registries.EncodingType) InlineEncodingParam {
	return &textParameter{
		ParamName: registries.Encoding,
		Value:     types.NewTextValue(string(value)),
	}
}
