package registry

// ActionRegistry contains all action names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.10
type ActionRegistry string

const (
	// Audio [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
	Audio ActionRegistry = "AUDIO"
	// Display [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
	Display ActionRegistry = "DISPLAY"
	// Email [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
	Email ActionRegistry = "EMAIL"
	// Procedure [See RFC-5545 spec]
	// Deprecated: As stated in the [RFC-5545], this registry value is deprecated
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
	// [RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
	Procedure ActionRegistry = "PROCEDURE"
)
