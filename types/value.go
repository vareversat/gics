package types

import "github.com/vareversat/gics/registry"

type V interface {
	GetIANAToken() string
}

type value struct {
	IANAToken registry.ValueTypeRegistry
}

func NewValue(IANAToken registry.ValueTypeRegistry) V {
	return &value{IANAToken: IANAToken}
}

func (v *value) GetIANAToken() string {
	return string(v.IANAToken)
}
