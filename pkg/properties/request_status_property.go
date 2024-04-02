package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.3

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type RequestStatusProperty interface {
	RequestStatusPropertyType
}

func NewRequestStatusProperty(
	code string,
	description string,
	extraData string,
	params ...parameters.Parameter,
) RequestStatusProperty {
	return &requestStatusPropertyType{
		PropName:          registries.REQUESTSTATUS,
		StatusCode:        types.NewTextValue(code),
		StatusDescription: types.NewTextValue(description),
		ExtraData:         types.NewTextValue(description),
		Parameters:        params,
	}
}
