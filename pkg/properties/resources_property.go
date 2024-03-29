package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.10

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type ResourcesProperty interface{}

type resourcesProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewResourcesProperty(value string) ResourcesProperty {
	return &resourcesProperty{
		IANAToken: registries.RESOURCES,
		Value:     values.NewTextValue(value),
	}
}
