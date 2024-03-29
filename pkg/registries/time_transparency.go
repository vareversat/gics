package registries

// RFC5545 section https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7

type TimeTransparency string

const (
	OPAQUE      TimeTransparency = "OPAQUE"
	TRANSPARENT TimeTransparency = "TRANSPARENT"
)
