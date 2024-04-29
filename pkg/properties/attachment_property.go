package properties

import (
	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type AttachmentProperty interface {
	TextPropertyType
}

// NewAttachmentProperty create a new ATTACH property
// This property CAN be seen in VALARM, VEVENT, VTODO & VJOURNAL components
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1
func NewAttachmentProperty(
	value string,
	params ...parameters.Parameter,
) AttachmentProperty {
	return &textPropertyType{
		PropName:   registries.ATTACH,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
