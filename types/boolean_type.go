package types

import (
	"github.com/vareversat/gics/registries"
)

type BooleanType interface {
	ValueType
	GetValue() bool
}

type booleanType struct {
	typeName  registries.ValueTypeRegistry
	typeValue bool
}

// GetValue get the bool typed value
func (bT *booleanType) GetValue() bool {
	return bT.typeValue
}

func (bT *booleanType) GetStringValue() string {
	if bT.typeValue {
		return "TRUE"
	} else {
		return "FALSE"
	}
}
func (bT *booleanType) GetTypeName() string {
	return string(bT.typeName)
}

// NewBooleanValue create a new [registries.Boolean] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.2
func NewBooleanValue(value bool) BooleanType {
	return &booleanType{
		typeName:  registries.Boolean,
		typeValue: value,
	}
}
