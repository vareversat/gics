package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.10

import (
	"strings"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type ResourcesProperty interface {
	TextPropertyType
}

func NewResourcesProperty(values []string, params ...parameters.Parameter) ResourcesProperty {
	return &textPropertyType{
		PropName:   registry.RESOURCES,
		Values:     types.NewTextValues(values),
		Parameters: params,
	}
}

func NewResourcesPropertyFromString(
	values string,
	params ...parameters.Parameter,
) ResourcesProperty {
	resources := strings.Split(values, ",")
	return &textPropertyType{
		PropName:   registry.RESOURCES,
		Values:     types.NewTextValues(resources),
		Parameters: params,
	}
}
