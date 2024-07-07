package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.3

import (
	"fmt"
	"strings"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
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
		ExtraData:         types.NewTextValue(extraData),
		Parameters:        params,
	}
}

func NewRequestStatusPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (RequestStatusProperty, error) {
	status := strings.Split(value, ";")
	if len(status) < 2 {
		return nil, fmt.Errorf(
			"the REQUEST-STATUS property is not formatted as CODE;DESCRIPTION[;EXTRA DATA]",
		)
	}
	code := status[0]
	description := status[1]
	extraData := ""

	if len(status) == 3 {
		extraData = status[1]
	}
	return &requestStatusPropertyType{
		PropName:          registries.REQUESTSTATUS,
		StatusCode:        types.NewTextValue(code),
		StatusDescription: types.NewTextValue(description),
		ExtraData:         types.NewTextValue(extraData),
		Parameters:        params,
	}, nil
}
