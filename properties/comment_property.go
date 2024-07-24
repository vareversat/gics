package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.4

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type CommentProperty interface {
	TextPropertyType
}

func NewCommentProperty(value string, params ...parameters.Parameter) CommentProperty {
	return &textPropertyType{
		PropName:   registries.Comment,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
