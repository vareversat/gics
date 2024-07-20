package types

import "github.com/vareversat/gics/registries"

type V interface {
	GetIANAToken() string
}

type value struct {
	IANAToken registries.ValueTypeRegistry
}

func NewValue(IANAToken registries.ValueTypeRegistry) V {
	return &value{IANAToken: IANAToken}
}

func (v *value) GetIANAToken() string {
	return string(v.IANAToken)
}
