package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type CalendarUserTypeParam interface {
	TextParameter
}

func NewCalendarUserTypeParam(value CalendarUserType) CalendarUserTypeParam {
	return &textParameter{
		ParamName: registries.CUTYPE,
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
