package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type CalendarUserTypeParam interface {
	TextTypeParameter
}

func NewCalendarUserTypeParam(value CalendarUserType) CalendarUserTypeParam {
	return &textParameter{
		ParamName: registry.CUTYPE,
		Value:     types.NewTextValue(string(value)),
	}
}

type CalendarUserType string

const (
	GROUP      CalendarUserType = "GROUP"
	INDIVIDUAL CalendarUserType = "INDIVIDUAL"
	RESOURCE   CalendarUserType = "RESOURCE"
	ROOM       CalendarUserType = "ROOM"
	UNKNOWN    CalendarUserType = "UNKNOWN"
)
