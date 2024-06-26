package registries

// RFC5545 section  https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.2

type PropertyNames string

const (
	ACTION          PropertyNames = "ACTION"
	ATTACH          PropertyNames = "ATTACH"
	ATTENDEE        PropertyNames = "ATTENDEE"
	BEGIN           PropertyNames = "BEGIN"
	CATEGORIES      PropertyNames = "CATEGORIES"
	CALSCALE        PropertyNames = "CALSCALE"
	CLASS           PropertyNames = "CLASS"
	COMMENT         PropertyNames = "COMMENT"
	COMPLETED_PROP  PropertyNames = "COMPLETED"
	CONTACT         PropertyNames = "CONTACT"
	CREATED         PropertyNames = "CREATED"
	DESCRIPTION     PropertyNames = "DESCRIPTION"
	DTEND           PropertyNames = "DTEND"
	DTSTART         PropertyNames = "DTSTART"
	DTSTAMP         PropertyNames = "DTSTAMP"
	DURATION_PROP   PropertyNames = "DURATION"
	DUE             PropertyNames = "DUE"
	END             PropertyNames = "END"
	EXDATE          PropertyNames = "EXDATE"
	EXRULE          PropertyNames = "EXRULE"
	FREEBUSY        PropertyNames = "FREEBUSY"
	GEO             PropertyNames = "GEO"
	LASTMODIFIED    PropertyNames = "LAST-MODIFIED"
	LOCATION        PropertyNames = "LOCATION"
	METHOD          PropertyNames = "METHOD"
	ORGANIZER       PropertyNames = "ORGANIZER"
	PERCENTCOMPLETE PropertyNames = "PERCENT-COMPLETE"
	PRIORITY        PropertyNames = "PRIORITY"
	PROP_TZID       PropertyNames = "TZID"
	PRODID          PropertyNames = "PRODID"
	RDATE           PropertyNames = "RDATE"
	RECURRENCEID    PropertyNames = "RECURRENCE-ID"
	RELATEDTO       PropertyNames = "RELATED-TO"
	REPEAT          PropertyNames = "REPEAT"
	REQUESTSTATUS   PropertyNames = "REQUEST-STATUS"
	RESOURCES       PropertyNames = "RESOURCES"
	RRULE           PropertyNames = "RRULE"
	SEQUENCE        PropertyNames = "SEQUENCE"
	STATUS          PropertyNames = "STATUS"
	SUMMARY         PropertyNames = "SUMMARY"
	TRANSP          PropertyNames = "TRANSP"
	TRIGGER         PropertyNames = "TRIGGER"
	TZNAME          PropertyNames = "TZNAME"
	TZOFFSETFROM    PropertyNames = "TZOFFSETFROM"
	TZOFFSETTO      PropertyNames = "TZOFFSETTO"
	TZURL           PropertyNames = "TZURL"
	UID             PropertyNames = "UID"
	URL             PropertyNames = "URL"
	VERSION         PropertyNames = "VERSION"
)
