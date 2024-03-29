package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.1

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type AlternateTextRepresentationParam interface{}

type alternateTextRepresentationParam struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewAlternateTextRepresentationParam(value string) AlternateTextRepresentationParam {
	return &alternateTextRepresentationParam{
		IANAToken: registries.ALTREP,
		Value:     value,
	}
}
