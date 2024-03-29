package registries

// RFC5545 section  https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.6

type TimeType string

const (
	FREE            TimeType = "FREE"
	BUSSY           TimeType = "BUSSY"
	BUSYUNAVAILABLE TimeType = "BUSY-UNAVAILABLE"
	BUSYTENTATIVE   TimeType = "BUSY-TENTATIVE"
)
