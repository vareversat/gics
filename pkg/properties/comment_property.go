package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.4

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type CommentProperty interface {
	TextPropertyType
}

func NewCommentProperty(value string, params ...parameters.Parameter) CommentProperty {
	return &textPropertyType{
		PropName:   registries.COMMENT,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
