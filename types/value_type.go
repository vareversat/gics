package types

// ValueType is the common interface of all typed value. See [RFC-5545] for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#autoid-32
type ValueType interface {
	// GetTypeName return the type name
	GetTypeName() string
	// GetStringValue return the string representation of the typed value
	GetStringValue() string
}
