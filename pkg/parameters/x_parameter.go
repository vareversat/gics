package parameters

import (
	"fmt"

	"github.com/vareversat/gics/pkg/types"
)

type XParameter interface{}

type xParameter struct {
	IANAToken string
	Value     types.TextValue
}

func NewXParameter(ianaToken string, value string) XParameter {
	return &xParameter{
		IANAToken: fmt.Sprintf("X-%s", ianaToken),
		Value:     types.NewTextValue(value),
	}
}
