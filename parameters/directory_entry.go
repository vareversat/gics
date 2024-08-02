package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DirectoryEntryParam interface {
	UriTypeParameter
}

// NewDirectoryEntryParam create a new registries.DirectoryEntryReference property
// This parameter can be used in these properties :
// - registries.Attendee
// - registries.Organizer
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6
func NewDirectoryEntryParam(value *url.URL) DirectoryEntryParam {
	return &uriParameter{
		ParamName: registries.DirectoryEntryReference,
		Value:     types.NewUriValue(value),
	}
}
