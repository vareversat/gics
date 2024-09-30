package properties

import (
	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AttachmentProperty interface {
	TextPropertyType
}

// NewAttachmentProperty create a new registries.AttachmentProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// Optional parameters :
// - registries.FormatTypeParam
// - registries.FormatTypeParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1
func NewAttachmentProperty(
	value string,
	params ...parameters.Parameter,
) AttachmentProperty {
	return &textPropertyType{
		PropName:   registries.AttachmentProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
