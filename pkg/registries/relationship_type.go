package registries

// RFC5545 section  https://https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.8

type RelationshipType string

const (
	CHILD   RelationshipType = "CHILD"
	PARENT  RelationshipType = "PARENT"
	SIBLING RelationshipType = "SIBLING"
)
