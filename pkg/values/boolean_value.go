package values

import (
	"github.com/vareversat/gics/pkg/registries"
)

type BinaryValue interface{}

type binaryValue struct {
	Base64Value string
	IANAToken   registries.Type
}

func NewBinaryValue(value string) BinaryValue {
	return &binaryValue{
		IANAToken:   registries.BINARY,
		Base64Value: value,
	}
}
