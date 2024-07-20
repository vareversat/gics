package registry

// FreeBusyTimeTypeRegistry contains all calendar user type registry names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.6
type FreeBusyTimeTypeRegistry string

const (
	// Free [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9
	Free FreeBusyTimeTypeRegistry = "FREE"
	// Busy [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9
	Busy FreeBusyTimeTypeRegistry = "BUSY"
	// BusyUnavailable [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9
	BusyUnavailable FreeBusyTimeTypeRegistry = "BUSY-UNAVAILABLE"
	// BusyTentative [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9
	BusyTentative FreeBusyTimeTypeRegistry = "BUSY-TENTATIVE"
)
