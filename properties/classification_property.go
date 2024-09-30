package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ClassificationProperty interface {
	TextPropertyType
}

// NewClassificationProperty create a new registries.ClassProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
func NewClassificationProperty(
	classValue registries.ClassificationRegistry,
	params ...parameters.Parameter,
) ClassificationProperty {
	return &classificationPropertyType{
		PropName:   registries.ClassProp,
		Value:      types.NewClassificationValue(classValue),
		Parameters: params,
	}
}

// NewClassificationPropertyFromString create a new registries.ClassProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
func NewClassificationPropertyFromString(
	classValue string,
	params ...parameters.Parameter,
) ClassificationProperty {
	return &classificationPropertyType{
		PropName:   registries.ClassProp,
		Value:      types.NewClassificationValue(registries.ClassificationRegistry(classValue)),
		Parameters: params,
	}
}
