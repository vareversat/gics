package values

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
)

type CalendarUserAddressValue interface{}

type calendarUserAddressValue struct {
	Value     url.URL
	IANAToken registries.Type
}

func NewCalendarUserAddressValue(value url.URL) CalendarUserAddressValue {
	return &calendarUserAddressValue{
		IANAToken: registries.CALADDRESS,
		Value:     value,
	}
}
