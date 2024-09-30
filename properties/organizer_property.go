package properties

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type OrganizerProperty interface {
	CalendarUserAddressPropertyType
}

// NewOrganizerProperty create a new registries.OrganizerProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Mandatory if group-scheduled calendar)
// - registries.Vtodo (Mandatory if group-scheduled calendar)
// - registries.Vjournal (Mandatory if group-scheduled calendar)
// - registries.Vfreebusy (Mandatory)
// Optional parameters :
// - registries.CommonNameParam
// - registries.DirectoryEntryReferenceParam
// - registries.SentByParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3
func NewOrganizerProperty(uri *url.URL, params ...parameters.Parameter) OrganizerProperty {
	return &calendarUserAddressPropertyType{
		PropName:   registries.OrganizerProp,
		Value:      types.NewCalendarUserAddressValue(uri),
		Parameters: params,
	}
}

// NewOrganizerPropertyFromString create a new registries.OrganizerProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Mandatory if group-scheduled calendar)
// - registries.Vtodo (Mandatory if group-scheduled calendar)
// - registries.Vjournal (Mandatory if group-scheduled calendar)
// - registries.Vfreebusy (Mandatory)
// Optional parameters :
// - registries.CommonNameParam
// - registries.DirectoryEntryReferenceParam
// - registries.SentByParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3
func NewOrganizerPropertyFromString(
	uri string,
	params ...parameters.Parameter,
) (AttendeeProperty, error) {
	urlValue, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", uri)
	}

	return &calendarUserAddressPropertyType{
		PropName:   registries.OrganizerProp,
		Value:      types.NewCalendarUserAddressValue(urlValue),
		Parameters: params,
	}, nil
}
