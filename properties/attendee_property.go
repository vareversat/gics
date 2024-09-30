package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AttendeeProperty interface {
	CalendarUserAddressPropertyType
}

// NewAttendeeProperty create a new registries.AttendeeProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// - registries.Vevent (Optional)
// - registries.Vfreebusy (Optional)
// - registries.Vjournal (Optional)
// Optional parameters :
// - registries.MemberParam
// - registries.CalendarUserTypeParam
// - registries.CommonNameParam
// - registries.DelegatedToParam
// - registries.DelegatedFromParam
// - registries.DirectoryEntryReferenceParam
// - registries.MemberParam
// - registries.RoleParam
// - registries.ParticipationStatusParam
// - registries.RsvpExpectationParam
// - registries.SentByParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
func NewAttendeeProperty(uri *url.URL, params ...parameters.Parameter) AttendeeProperty {
	return &calendarUserAddressPropertyType{
		PropName:   registries.AttendeeProp,
		Value:      types.NewCalendarUserAddressValue(uri),
		Parameters: params,
	}
}

// NewAttendeePropertyFromString create a new registries.AttendeeProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// - registries.Vevent (Optional)
// - registries.Vfreebusy (Optional)
// - registries.Vjournal (Optional)
// Optional parameters :
// - registries.MemberParam
// - registries.CalendarUserTypeParam
// - registries.CommonNameParam
// - registries.DelegatedToParam
// - registries.DelegatedFromParam
// - registries.DirectoryEntryReferenceParam
// - registries.MemberParam
// - registries.RoleParam
// - registries.ParticipationStatusParam
// - registries.RsvpExpectationParam
// - registries.SentByParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
func NewAttendeePropertyFromString(
	uri string,
	params ...parameters.Parameter,
) (OrganizerProperty, error) {
	urlValue, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", uri)
	}

	return &calendarUserAddressPropertyType{
		PropName:   registries.AttendeeProp,
		Value:      types.NewCalendarUserAddressValue(urlValue),
		Parameters: params,
	}, nil
}
