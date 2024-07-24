package properties

import (
	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AttachmentProperty interface {
	TextPropertyType
}

// NewAttachmentProperty create a new ATTACH property
// This property CAN be seen in VALARM, VEVENT, VTODO & VJOURNAL components
// See the [RFC-5545] ref for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1
func NewAttachmentProperty(
	value string,
	params ...parameters.Parameter,
) AttachmentProperty {
	return &textPropertyType{
		PropName:   registries.Attachment,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
