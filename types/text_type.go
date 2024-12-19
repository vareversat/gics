package types

import "github.com/vareversat/gics/registries"

type TextType interface {
	ValueType
}

type textType struct {
	typeName  registries.ValueTypeRegistry
	typeValue string
}

func (t *textType) GetStringValue() string {
	return t.typeValue
}

func (t *textType) GetTypeName() string {
	return string(t.typeName)
}

// NewTextValue create a new [registries.Text] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.11
func NewTextValue(value string) TextType {
	return &textType{
		typeName:  registries.Text,
		typeValue: value,
	}
}

// NewTextValues create a new [registries.Text] type values. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.11
func NewTextValues(values []string) []TextType {
	textValues := make([]TextType, 0)

	for i := 0; i < len(values); i++ {
		textValues = append(textValues, NewTextValue(values[i]))
	}
	return textValues
}
