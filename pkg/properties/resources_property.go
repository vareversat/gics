package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.10

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type ResourcesProperty interface {
	TextPropertyType
}

func NewResourcesProperty(values []string, params ...parameters.Parameter) ResourcesProperty {
	return &textPropertyType{
		PropName:   registries.RESOURCES,
		Values:     types.NewTextValues(values),
		Parameters: params,
	}
}
