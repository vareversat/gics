package types

import "github.com/vareversat/gics/registries"

type StatusType interface {
	ValueType
	GetValue() registries.ParticipantStatusRegistry
}

type statusType struct {
	typeName  registries.ValueTypeRegistry
	typeValue registries.ParticipantStatusRegistry
}

func (s *statusType) GetStringValue() string {
	return string(s.typeValue)
}

func (s *statusType) GetTypeName() string {
	return string(s.typeName)
}

// GetValue get the [registries.ParticipantStatusRegistry] typed value
func (s *statusType) GetValue() registries.ParticipantStatusRegistry {
	return s.typeValue
}

func NewStatusValue(value registries.ParticipantStatusRegistry) StatusType {
	return &statusType{
		typeName:  registries.Text,
		typeValue: value,
	}
}
