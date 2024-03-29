package registries

// RFC5545 section https://https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.11

type Classification string

const (
	PUBLIC       Classification = "PUBLIC"
	PRIVATE      Classification = "PRIVATE"
	CONFIDENTIAL Classification = "CONFIDENTIAL"
)
