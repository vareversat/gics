package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1

import (
	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type AttachmentProperty interface {
	TextPropertyType
}

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
