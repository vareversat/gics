package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.7

// Parameter used in these properties :
// - BINARY

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type InlineEncodingParam interface {
	TextParameter
}

func NewInlineEncodingParam(value EncodingType) InlineEncodingParam {
	return &textParameter{
		ParamName: registries.ENCODING,
		Value:     types.NewTextValue(string(value)),
	}
}

type EncodingType string

const (
	_8BIT  EncodingType = "8BIT"
	BASE64 EncodingType = "BASE64"
)
