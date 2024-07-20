package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.7

// Parameter used in these properties :
// - BINARY

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type InlineEncodingParam interface {
	TextTypeParameter
}

func NewInlineEncodingParam(value EncodingType) InlineEncodingParam {
	return &textParameter{
		ParamName: registry.ENCODING,
		Value:     types.NewTextValue(string(value)),
	}
}

type EncodingType string

const (
	_8BIT  EncodingType = "8BIT"
	BASE64 EncodingType = "BASE64"
)
