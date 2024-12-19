package properties

import (
	"fmt"
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
	long float64,
	lat float64,
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
	return &geoPropertyType{
		PropName:   registries.GeoProp,
		Longitude:  types.NewFloatValueFromString(coordinates[0]),
		Latitude:   types.NewFloatValueFromString(coordinates[1]),
		Parameters: params,
	}, nil
}
