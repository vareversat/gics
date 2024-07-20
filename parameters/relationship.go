package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10

// Parameter used in these properties :
// - RELATED-TO

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RelationshipParam interface {
	TextTypeParameter
}

func NewRelationshipParam(value registries.RelationshipTypeRegistry) RelationshipParam {
	return &textParameter{
		ParamName: registries.RelationshipType,
		Value:     types.NewTextValue(string(value)),
	}
}
