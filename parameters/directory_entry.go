package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6

// Parameter used in these properties :
// - ATTENDEE
// - ORGANIZER

type DirectoryEntryParam interface {
	UriTypeParameter
}

func NewDirectoryEntryParam(value *url.URL) DirectoryEntryParam {
	return &uriParameter{
		ParamName: registry.DirectoryEntryReference,
		Value:     types.NewUriValue(value),
	}
}
