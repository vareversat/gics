package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type GeographicPositionProperty interface {
	GeoPropertyType
}

func NewGeographicPositionProperty(
	long float32,
	lat float32,
	params ...parameters.Parameter,
) GeographicPositionProperty {
	return &geoPropertyType{
		PropName:   registries.GEO,
		Longitude:  types.NewFloatValue(long),
		Latitude:   types.NewFloatValue(lat),
		Parameters: params,
	}
}

func NewGeographicPositionPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (GeographicPositionProperty, error) {
	coordinates := strings.Split(value, ";")
	if len(coordinates) != 2 {
		return nil, fmt.Errorf("the GEO property is not formatted as LATITUDE;LONGITUDE")
	}
	long, err := strconv.ParseFloat(coordinates[0], 32)
	lat, err := strconv.ParseFloat(coordinates[1], 32)

	if err != nil {
		return nil, fmt.Errorf(
			"%s or %s cannot be parsed as float32",
			coordinates[0],
			coordinates[1],
		)
	}
	return &geoPropertyType{
		PropName:   registries.GEO,
		Longitude:  types.NewFloatValue(float32(long)),
		Latitude:   types.NewFloatValue(float32(lat)),
		Parameters: params,
	}, nil
}
