package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type VersionProperty interface {
}

type versionProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewVersionProperty(value string) VersionProperty {
	return &versionProperty{
		IANAToken: registries.VERSION,
		Value:     values.NewTextValue(value),
	}
}
