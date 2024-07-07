package properties

import (
	"fmt"

	"github.com/vareversat/gics/types"
)

type XProperty interface{}

type xProperty struct {
	PropName string
	Value    types.TextValue
}

func NewXProperty(PropName string, value string) XProperty {
	return &xProperty{
		PropName: fmt.Sprintf("X-%s", PropName),
		Value:    types.NewTextValue(value),
	}
}
