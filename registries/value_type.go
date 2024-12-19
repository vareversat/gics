package registries

// ValueTypeRegistry contains all value type names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.4
type ValueTypeRegistry string

const (
	// Binary. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.1
	Binary ValueTypeRegistry = "BINARY"
	// Boolean. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.2
	Boolean ValueTypeRegistry = "BOOLEAN"
	// CalAddress. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.3
	CalAddress ValueTypeRegistry = "CAL-ADDRESS"
	// Date. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4
	Date ValueTypeRegistry = "DATE"
	// DateTime. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.5
	DateTime ValueTypeRegistry = "DATE-TIME"
	// Duration. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
	Duration ValueTypeRegistry = "DURATION"
	// Float. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.7
	Float ValueTypeRegistry = "FLOAT"
	// Integer. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.8
	Integer ValueTypeRegistry = "INTEGER"
	// PeriodOfTime. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9
	PeriodOfTime ValueTypeRegistry = "PERIOD"
	// Recur. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
	Recur ValueTypeRegistry = "RECUR"
	// Text. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.11
	Text ValueTypeRegistry = "TEXT"
	// Uri. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.12
	Uri ValueTypeRegistry = "URI"
	// Time. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.13
	Time ValueTypeRegistry = "TIME"
	// UTCOffset. [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.14
	UTCOffset ValueTypeRegistry = "UTC-OFFSET"
)
