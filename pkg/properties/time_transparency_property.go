package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type TimeTransparencyProperty interface {
}

type timeTransparencyProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewTimeTransparencyProperty(value registries.TimeTransparency) TimeTransparencyProperty {
	return &timeTransparencyProperty{
		IANAToken: registries.TRANSP,
		Value:     values.NewTextValue(string(value)),
	}
}
