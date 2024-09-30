package properties

import (
	"fmt"
	"strings"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RequestStatusProperty interface {
	RequestStatusPropertyType
}

// NewRequestStatusProperty create a new registries.RequestStatusProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vfreebusy (Optional)
// Optional parameters :
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.3
func NewRequestStatusProperty(
	code string,
	description string,
	extraData string,
	params ...parameters.Parameter,
) RequestStatusProperty {
	return &requestStatusPropertyType{
		PropName:          registries.RequestStatusProp,
		StatusCode:        types.NewTextValue(code),
		StatusDescription: types.NewTextValue(description),
		ExtraData:         types.NewTextValue(extraData),
		Parameters:        params,
	}
}

// NewRequestStatusPropertyFromString create a new registries.RequestStatusProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vfreebusy (Optional)
// Optional parameters :
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.3
func NewRequestStatusPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (RequestStatusProperty, error) {
	status := strings.Split(value, ";")
	if len(status) < 2 {
		return nil, fmt.Errorf(
			"the REQUEST-STATUS property is not formatted as CODE;DESCRIPTION[;EXTRA DATA]",
		)
	}
	code := status[0]
	description := status[1]
	extraData := ""

	if len(status) == 3 {
		extraData = status[1]
	}
	return &requestStatusPropertyType{
		PropName:          registries.RequestStatusProp,
		StatusCode:        types.NewTextValue(code),
		StatusDescription: types.NewTextValue(description),
		ExtraData:         types.NewTextValue(extraData),
		Parameters:        params,
	}, nil
}
