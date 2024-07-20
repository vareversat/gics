package registry

// RFC5545 section https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.4

type Type string

const (
	BINARY     Type = "BINARY"
	BOOLEAN    Type = "BOOLEAN"
	CALADDRESS Type = "CAL-ADDRESS"
	DATE       Type = "DATE"
	DATETIME   Type = "DATE-TIME"
	DURATION   Type = "DURATION"
	FLOAT      Type = "FLOAT"
	INTEGER    Type = "INTEGER"
	PERIOD     Type = "PERIOD"
	RECUR      Type = "RECUR"
	TEXT       Type = "TEXT"
	URI        Type = "URI"
	TIME       Type = "TIME"
	UTCOFFSET  Type = "UTC-OFFSET"
)
