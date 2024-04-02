package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10

// Parameter used in these properties :
// - RELATED-TO

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type RelationshipParam interface {
	TextParameter
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
