package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RelationshipParam interface {
	TextTypeParameter
}

// NewRelationshipParam create a new registries.RelationshipTypeParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.RelatedToProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10
func NewRelationshipParam(value registries.RelationshipTypeRegistry) RelationshipParam {
	return &textParameter{
		ParamName: registries.RelationshipTypeParam,
		Value:     types.NewTextValue(string(value)),
	}
}

func NewRelationshipParamFromString(value string) RelationshipParam {
	return &textParameter{
		ParamName: registries.RelationshipTypeParam,
		Value:     types.NewTextValue(value),
	}
}
