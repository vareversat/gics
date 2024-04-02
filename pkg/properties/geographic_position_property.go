package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.6

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type GeographicPositionProperty interface {
	GeoPropertyType
}

func NewGeographicPositionProperty(long float32, lat float32) GeographicPositionProperty {
	return &geoPropertyType{
		PropName:  registries.GEO,
		Longitude: types.NewFloatValue(long),
		Latitude:  types.NewFloatValue(lat),
	}
}
