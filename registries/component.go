package registries

// ComponentRegistry contains all component names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.1
type ComponentRegistry string

const (
	// Vcalendar [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.4
	Vcalendar ComponentRegistry = "VCALENDAR"
	// Vevent [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.1
	Vevent ComponentRegistry = "VEVENT"
	// Vtodo [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.2
	Vtodo ComponentRegistry = "VTODO"
	// Vjournal [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.3
	Vjournal ComponentRegistry = "VJOURNAL"
	// Vfreebusy [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.4
	Vfreebusy ComponentRegistry = "VFREEBUSY"
	// Vtimezone [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.5
	Vtimezone ComponentRegistry = "VTIMEZONE"
	// Valarm [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6
	Valarm ComponentRegistry = "VALARM"
	// Standard [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.7
	Standard ComponentRegistry = "STANDARD"
	// Daylight [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.8
	Daylight ComponentRegistry = "DAYLIGHT"
)
