package registries

// RFC5545 section  https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.10

type Action string

const (
	AUDIO   TimeType = "AUDIO"
	DISPLAY TimeType = "DISPLAY"
	EMAIL   TimeType = "EMAIL"
)
