package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type AttachmentProperty interface{}

type attachmentProperty struct {
	IANAToken   registries.Properties
	Value       values.UriValue // Either Binary OR URI
	BinaryValue values.BinaryValue
}

func NewAttachmentProperty(value url.URL, binaryValue string) AttachmentProperty {
	return &attachmentProperty{
		IANAToken:   registries.ATTACH,
		Value:       values.NewUriValue(value),
		BinaryValue: values.NewBinaryValue(binaryValue),
	}
}
