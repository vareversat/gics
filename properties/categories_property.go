package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
	"github.com/vareversat/gics/utils"
)

type CategoriesProperty interface {
	TextPropertyType
}

// NewCategoryProperty create a new registries.CategoriesProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// Optional parameters :
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.2
func NewCategoryProperty(values []string, params ...parameters.Parameter) CategoriesProperty {
	return &textPropertyType{
		PropName:   registries.CategoriesProp,
		Values:     types.NewTextValues(values),
		Parameters: params,
	}
}

// NewCategoryPropertyFromString create a new registries.CategoriesProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// Optional parameters :
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.2
func NewCategoryPropertyFromString(
	values string,
	params ...parameters.Parameter,
) CategoriesProperty {
	return &textPropertyType{
		PropName:   registries.CategoriesProp,
		Values:     types.NewTextValues(utils.StringToStringArray(values)),
		Parameters: params,
	}
}
