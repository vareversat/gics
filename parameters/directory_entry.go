package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DirectoryEntryParam interface {
	UriTypeParameter
}

// NewDirectoryEntryParam create a new registries.DirectoryEntryReferenceParam property
// This parameter can be used in these properties :
// - registries.AttendeeProp
// - registries.OrganizerProp
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6
func NewDirectoryEntryParam(value *url.URL) DirectoryEntryParam {
	return &uriParameter{
		ParamName: registries.DirectoryEntryReferenceParam,
		Value:     types.NewUriValue(value),
	}
}
