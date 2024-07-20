package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10

// Parameter used in these properties :
// - RELATED-TO

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type RelationshipParam interface {
	TextTypeParameter
}

func NewRelationshipParam(value registry.RelationshipTypeRegistry) RelationshipParam {
	return &textParameter{
		ParamName: registry.RelationshipType,
		Value:     types.NewTextValue(string(value)),
	}
}
