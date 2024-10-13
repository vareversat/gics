package parameters

import (
	"github.com/vareversat/gics/types"
)

func NewNonStandardParameter(parameterName string, parameterValue string) NonStandardParameter {
	return &nonStandardParameter{
		ParamName: parameterName,
		Value:     types.NewTextValue(parameterValue),
	}
}
