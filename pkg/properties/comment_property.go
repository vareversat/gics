package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.4

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type CommentProperty interface{}

type commentProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewCommentProperty(value string) CommentProperty {
	return &commentProperty{
		IANAToken: registries.COMMENT,
		Value:     values.NewTextValue(value),
	}
}
