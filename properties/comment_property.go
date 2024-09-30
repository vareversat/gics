package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type CommentProperty interface {
	TextPropertyType
}

// NewCommentProperty create a new registries.CommentProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vtimezone (Optional)
// - registries.Vfreebusy (Optional)
// Optional parameters :
// - registries.AlternateTextRepresentationParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.4
func NewCommentProperty(value string, params ...parameters.Parameter) CommentProperty {
	return &textPropertyType{
		PropName:   registries.CommentProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
