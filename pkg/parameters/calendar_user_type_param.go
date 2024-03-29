package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type CalendarUserTypeParam interface{}

type calendarUserTypeParam struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewCalendarUserTypeParam(value registries.CalendarUserType) CalendarUserTypeParam {
	return &calendarUserTypeParam{
		IANAToken: registries.CUTYPE,
		Value:     values.NewTextValue(string(value)),
	}
}
