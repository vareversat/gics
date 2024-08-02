package registries

// PropertyRegistry contains all property names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.2
type PropertyRegistry string

const (
	// CalendarScale [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.1
	CalendarScale PropertyRegistry = "CALSCALE"
	// Method [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.2
	Method PropertyRegistry = "METHOD"
	// ProductIdentifier [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.3
	ProductIdentifier PropertyRegistry = "PRODID"
	// Version [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4
	Version PropertyRegistry = "VERSION"
	// Attachment [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1
	Attachment PropertyRegistry = "ATTACH"
	// Categories [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.2
	Categories PropertyRegistry = "CATEGORIES"
	// Class [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
	Class PropertyRegistry = "CLASS"
	// Comment [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.4
	Comment PropertyRegistry = "COMMENT"
	// Description [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5
	Description PropertyRegistry = "DESCRIPTION"
	// Geo [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.6
	Geo PropertyRegistry = "GEO"
	// Location [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.7
	Location PropertyRegistry = "LOCATION"
	// PercentComplete [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.8
	PercentComplete PropertyRegistry = "PERCENT-COMPLETE"
	// Priority [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.9
	Priority PropertyRegistry = "PRIORITY"
	// Resources [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.10
	Resources PropertyRegistry = "RESOURCES"
	// Status [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11
	Status PropertyRegistry = "STATUS"
	// Summary [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12
	Summary PropertyRegistry = "SUMMARY"
	// CompletedProperty [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.1
	CompletedProperty PropertyRegistry = "COMPLETED"
	// DateTimeEnd [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.2
	DateTimeEnd PropertyRegistry = "DTEND"
	// DateTimeDue [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.3
	DateTimeDue PropertyRegistry = "DUE"
	// DateTimeStart [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.4
	DateTimeStart PropertyRegistry = "DTSTART"
	// DurationProperty [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.5
	DurationProperty PropertyRegistry = "DURATION"
	// FreeBusyTime [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.6
	FreeBusyTime PropertyRegistry = "FREEBUSY"
	// TimeTransparency [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
	TimeTransparency PropertyRegistry = "TRANSP"
	// TimeZoneIdProperty [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.1
	TimeZoneIdProperty PropertyRegistry = "TZID"
	// TimeZoneName [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.2
	TimeZoneName PropertyRegistry = "TZNAME"
	// TimeZoneOffsetFrom [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.3
	TimeZoneOffsetFrom PropertyRegistry = "TZOFFSETFROM"
	// TimeZoneOffsetTo [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.4
	TimeZoneOffsetTo PropertyRegistry = "TZOFFSETTO"
	// TimeZoneUrl [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.5
	TimeZoneUrl PropertyRegistry = "TZURL"
	// Attendee [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
	Attendee PropertyRegistry = "ATTENDEE"
	// Contact [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2
	Contact PropertyRegistry = "CONTACT"
	// Organizer [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3
	Organizer PropertyRegistry = "ORGANIZER"
	// RecurrenceId [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.4
	RecurrenceId PropertyRegistry = "RECURRENCE-ID"
	// RelatedTo [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.5
	RelatedTo PropertyRegistry = "RELATED-TO"
	// Url [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.6
	Url PropertyRegistry = "URL"
	// Uid [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7
	Uid PropertyRegistry = "UID"
	// ExceptionDateTimes [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.1
	ExceptionDateTimes PropertyRegistry = "EXDATE"
	// ExceptionRule [See RFC-5545 spec]
	// Deprecated: As stated in the [RFC-5545], this registry value is deprecated
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
	// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
	ExceptionRule PropertyRegistry = "EXRULE"
	// RecurrenceDateTimes [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
	RecurrenceDateTimes PropertyRegistry = "RDATE"
	// RecurrenceRule [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.3
	RecurrenceRule PropertyRegistry = "RRULE"
	// Action [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
	Action PropertyRegistry = "ACTION"
	// Repeat [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.2
	Repeat PropertyRegistry = "REPEAT"
	// Trigger [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3
	Trigger PropertyRegistry = "TRIGGER"
	// DateTimeCreated [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.1
	DateTimeCreated PropertyRegistry = "CREATED"
	// DateTimeStamp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2
	DateTimeStamp PropertyRegistry = "DTSTAMP"
	// LastModified [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3
	LastModified PropertyRegistry = "LAST-MODIFIED"
	// Sequence [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4
	Sequence PropertyRegistry = "SEQUENCE"
	// RequestStatus [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.3
	RequestStatus PropertyRegistry = "REQUEST-STATUS"

	// Begin is used to start a new component definition
	Begin PropertyRegistry = "BEGIN"
	// End is used to end a component definition
	End PropertyRegistry = "END"
)
