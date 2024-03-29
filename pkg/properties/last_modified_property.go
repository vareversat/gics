package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type LastModifiedProperty interface {
}

type lastModifiedProperty struct {
	IANAToken registries.Properties
	Value     values.DateTimeValue
}

func NewLastModifiedProperty(value *time.Time) LastModifiedProperty {
	return &lastModifiedProperty{
		IANAToken: registries.LASTMODIFIED,
		Value:     values.NewDateTimeValue(value),
	}
}
