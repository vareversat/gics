package properties

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type GeographicPositionProperty interface {
	GeoPropertyType
}

// NewGeographicPositionProperty create a new registries.FreeBusyTimeProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.6
func NewGeographicPositionProperty(
	long float32,
	lat float32,
	params ...parameters.Parameter,
) GeographicPositionProperty {
	return &geoPropertyType{
		PropName:   registries.GeoProp,
		Longitude:  types.NewFloatValue(long),
		Latitude:   types.NewFloatValue(lat),
		Parameters: params,
	}
}

// NewGeographicPositionPropertyFromString create a new registries.FreeBusyTimeProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.6
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
		PropName:   registries.GeoProp,
		Longitude:  types.NewFloatValue(float32(long)),
		Latitude:   types.NewFloatValue(float32(lat)),
		Parameters: params,
	}, nil
}
