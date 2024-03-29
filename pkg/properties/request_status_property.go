package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type RequestStatusProperty interface{}

type requestStatusProperty struct {
	IANAToken         registries.Properties
	StatusCode        values.TextValue
	StatusDescription values.TextValue
}

func NewRequestStatusProperty(code string, description string) RequestStatusProperty {
	return &requestStatusProperty{
		IANAToken:         registries.REQUESTSTATUS,
		StatusCode:        values.NewTextValue(code),
		StatusDescription: values.NewTextValue(description),
	}
}
