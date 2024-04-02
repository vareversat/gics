package registries

// RFC5545 section  https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.3

type Parameters string

const (
	ALTREP        Parameters = "ALTREP"
	CN            Parameters = "CN"
	CUTYPE        Parameters = "CUTYPE"
	DELEGATEDFROM Parameters = "DELEGATED-FROM"
	DELEGATEDTO   Parameters = "DELEGATED-TO"
	DIR           Parameters = "DIR"
	ENCODING      Parameters = "ENCODING"
	FMTTYPE       Parameters = "FMTTYPE"
	FBTYPE        Parameters = "FBTYPE"
	LANGUAGE      Parameters = "LANGUAGE"
	MEMBER        Parameters = "MEMBER"
	PARTSTAT      Parameters = "PARTSTAT"
	RANGE         Parameters = "RANGE"
	RELATED       Parameters = "RELATED"
	RELTYPE       Parameters = "RELTYPE"
	ROLE          Parameters = "ROLE"
	RSVP          Parameters = "RSVP"
	SENTBY        Parameters = "SENT-BY"
	PARAM_TZID    Parameters = "TZID"
	VALUE         Parameters = "VALUE"
)
