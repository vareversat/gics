package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type CommonNameParam interface{}

type commonNameParam struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewCommonNameParam(value string) CommonNameParam {
	return &commonNameParam{
		IANAToken: registries.CN,
		Value:     value,
	}
}
