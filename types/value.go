package types

import "github.com/vareversat/gics/registries"

type V interface {
	GetIANAToken() string
}

type value struct {
	IANAToken registries.Type
}

func NewValue(IANAToken registries.Type) V {
	return &value{IANAToken: IANAToken}
}

func (v *value) GetIANAToken() string {
	return string(v.IANAToken)
}
