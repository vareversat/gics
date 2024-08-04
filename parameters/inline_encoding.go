package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type InlineEncodingParam interface {
	TextTypeParameter
}

// NewInlineEncodingParam create a new registries.EncodingParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.AttachmentProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.7
func NewInlineEncodingParam(value registries.EncodingTypeRegistry) InlineEncodingParam {
	return &textParameter{
		ParamName: registries.EncodingParam,
		Value:     types.NewTextValue(string(value)),
	}
}
