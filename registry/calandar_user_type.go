package registry

// CalendarUserTypeRegistry contains all calendar user type registry names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.5
type CalendarUserTypeRegistry string

const (
	// Group [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	Group CalendarUserTypeRegistry = "GROUP"
	// Individual [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	Individual CalendarUserTypeRegistry = "INDIVIDUAL"
	// Resource [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	Resource CalendarUserTypeRegistry = "RESOURCE"
	// Room [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	Room CalendarUserTypeRegistry = "ROOM"
	// Unknown [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	Unknown CalendarUserTypeRegistry = "UNKNOWN"
	// Group [See RFC-5545 spec]
)
