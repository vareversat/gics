package types

import (
	"github.com/vareversat/gics/registries"
)

type BinaryType interface {
	ValueType
	GetValue() string
}

type binaryType struct {
	typeName  registries.ValueTypeRegistry
	typeValue string
}

// GetValue get the string (base64) typed value
func (bT *binaryType) GetValue() string {
	return string(bT.typeValue)
}

func (bT *binaryType) GetStringValue() string {
	return string(bT.typeValue)
}

func (bT *binaryType) GetTypeName() string {
	return string(bT.typeName)
}

// NewBinaryValue create a new [registries.Binary] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.1
func NewBinaryValue(value string) BinaryType {
	return &binaryType{
		typeName:  registries.Binary,
		typeValue: value,
	}
}
