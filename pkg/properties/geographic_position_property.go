package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.6

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type GeographicPositionProperty interface{}

type geographicPositionProperty struct {
	IANAToken registries.Properties
	Longitude values.FloatValue
	Latitude  values.FloatValue
}

func NewGeographicPositionProperty(long float32, lat float32) GeographicPositionProperty {
	return &geographicPositionProperty{
		IANAToken: registries.GEO,
		Longitude: values.NewFloatValue(long),
		Latitude:  values.NewFloatValue(lat),
	}
}
