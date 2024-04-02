package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type VersionProperty interface {
	TextPropertyType
}

func NewVersionProperty(value string) VersionProperty {
	return &textPropertyType{
		PropName: registries.VERSION,
		Value:    types.NewTextValue(value),
	}
}
