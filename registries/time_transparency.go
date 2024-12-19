package registries

// TimeTransparencyRegistry contains all transparency type registry names of the [RFC-5545]
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
type TimeTransparencyRegistry string

const (
	// Opaque [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
	Opaque TimeTransparencyRegistry = "OPAQUE"
	// Transparent [See RFC-5545 spec]
	//
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
	Transparent TimeTransparencyRegistry = "TRANSPARENT"
)
