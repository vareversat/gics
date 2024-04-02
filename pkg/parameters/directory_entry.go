package parameters

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6

// Parameter used in these properties :
// - ATTENDEE
// - ORGANIZER

type DirectoryEntryParam interface {
	UriParameter
}

func NewDirectoryEntryParam(value *url.URL) DirectoryEntryParam {
	return &uriParameter{
		ParamName: registries.DIR,
		Value:     types.NewUriValue(value),
	}
}
