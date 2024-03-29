package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type ContactProperty interface{}

type contactProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewContactProperty(value string) ContactProperty {
	return &contactProperty{
		IANAToken: registries.CONTACT,
		Value:     values.NewTextValue(value),
	}
}
