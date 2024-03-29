package registries

// RFC5545 section  https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.3

type Parameters string

const (
	ALTREP        Properties = "ALTREP"
	CN            Properties = "CN"
	CUTYPE        Properties = "CUTYPE"
	DELEGATEDFROM Properties = "DELEGATED-FROM"
	DELEGATEDTO   Properties = "DELEGATED-TO"
	DIR           Properties = "DIR"
	ENCODING      Properties = "ENCODING"
	FMTTYPE       Properties = "FMTTYPE"
	FBTYPE        Properties = "FBTYPE"
	LANGUAGE      Properties = "LANGUAGE"
	MEMBER        Properties = "MEMBER"
	PARTSTAT      Properties = "PARTSTAT"
	RANGE         Properties = "RANGE"
	RELATED       Properties = "RELATED"
	RELTYPE       Properties = "RELTYPE"
	ROLE          Properties = "ROLE"
	RSVP          Properties = "RSVP"
	SENTBY        Properties = "SENT-BY"
	PARAM_TZID    Properties = "TZID"
	VALUE         Properties = "VALUE"
)
