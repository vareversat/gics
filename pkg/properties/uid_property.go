package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type UidProperty interface {
}

type uidProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewUidProperty(value string) UidProperty {
	return &uidProperty{
		IANAToken: registries.VERSION,
		Value:     values.NewTextValue(value),
	}
}
