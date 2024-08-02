package registries

// EncodingType contains all relationship names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.8
type EncodingType string

const (
	// _8BIT [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.8
	_8BIT EncodingType = "8BIT"
	// Base64 [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.8
	Base64 EncodingType = "BASE64"
)
