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

func NewRelationshipParam(value RelationshipType) RelationshipParam {
	return &textParameter{
		ParamName: registries.RELTYPE,
		Value:     types.NewTextValue(string(value)),
	}
}

type RelationshipType string

const (
	CHILD   RelationshipType = "CHILD"
	PARENT  RelationshipType = "PARENT"
	SIBLING RelationshipType = "SIBLING"
)
