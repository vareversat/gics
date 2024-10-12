package registries

// PropertyRegistry contains all property names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.2
type PropertyRegistry string

const (
	// CalendarScaleProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.1
	CalendarScaleProp PropertyRegistry = "CALSCALE"
	// MethodProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.2
	MethodProp PropertyRegistry = "METHOD"
	// ProductIdentifierProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.3
	ProductIdentifierProp PropertyRegistry = "PRODID"
	// VersionProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4
	VersionProp PropertyRegistry = "VERSION"
	// AttachmentProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1
	AttachmentProp PropertyRegistry = "ATTACH"
	// CategoriesProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.2
	CategoriesProp PropertyRegistry = "CATEGORIES"
	// ClassProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
	ClassProp PropertyRegistry = "CLASS"
	// CommentProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.4
	CommentProp PropertyRegistry = "COMMENT"
	// DescriptionProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5
	DescriptionProp PropertyRegistry = "DESCRIPTION"
	// GeoProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.6
	GeoProp PropertyRegistry = "GEO"
	// LocationProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.7
	LocationProp PropertyRegistry = "LOCATION"
	// PercentCompleteProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.8
	PercentCompleteProp PropertyRegistry = "PERCENT-COMPLETE"
	// PriorityProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.9
	PriorityProp PropertyRegistry = "PRIORITY"
	// ResourcesProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.10
	ResourcesProp PropertyRegistry = "RESOURCES"
	// StatusProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11
	StatusProp PropertyRegistry = "STATUS"
	// SummaryProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12
	SummaryProp PropertyRegistry = "SUMMARY"
	// CompletedProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.1
	CompletedProp PropertyRegistry = "COMPLETED"
	// DateTimeEndProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.2
	DateTimeEndProp PropertyRegistry = "DTEND"
	// DateTimeDueProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.3
	DateTimeDueProp PropertyRegistry = "DUE"
	// DateTimeStartProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.4
	DateTimeStartProp PropertyRegistry = "DTSTART"
	// DurationProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.5
	DurationProp PropertyRegistry = "DURATION"
	// FreeBusyTimeProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.6
	FreeBusyTimeProp PropertyRegistry = "FREEBUSY"
	// TimeTransparencyProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
	TimeTransparencyProp PropertyRegistry = "TRANSP"
	// TimeZoneIdProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.1
	TimeZoneIdProp PropertyRegistry = "TZID"
	// TimeZoneNameProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.2
	TimeZoneNameProp PropertyRegistry = "TZNAME"
	// TimeZoneOffsetFromProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.3
	TimeZoneOffsetFromProp PropertyRegistry = "TZOFFSETFROM"
	// TimeZoneOffsetToProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.4
	TimeZoneOffsetToProp PropertyRegistry = "TZOFFSETTO"
	// TimeZoneUrlProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.5
	TimeZoneUrlProp PropertyRegistry = "TZURL"
	// AttendeeProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
	AttendeeProp PropertyRegistry = "ATTENDEE"
	// ContactProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2
	ContactProp PropertyRegistry = "CONTACT"
	// OrganizerProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3
	OrganizerProp PropertyRegistry = "ORGANIZER"
	// RecurrenceIdProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.4
	RecurrenceIdProp PropertyRegistry = "RECURRENCE-ID"
	// RelatedToProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.5
	RelatedToProp PropertyRegistry = "RELATED-TO"
	// UrlProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.6
	UrlProp PropertyRegistry = "URL"
	// UidProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7
	UidProp PropertyRegistry = "UID"
	// ExceptionDateTimesProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.1
	ExceptionDateTimesProp PropertyRegistry = "EXDATE"
	// ExceptionRuleProp [See RFC-5545 spec]
	// Deprecated: As stated in the [RFC-5545], this registry value is deprecated
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
	// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
	ExceptionRuleProp PropertyRegistry = "EXRULE"
	// RecurrenceDateTimesProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
	RecurrenceDateTimesProp PropertyRegistry = "RDATE"
	// RecurrenceRuleProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.3
	RecurrenceRuleProp PropertyRegistry = "RRULE"
	// ActionProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
	ActionProp PropertyRegistry = "ACTION"
	// RepeatProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.2
	RepeatProp PropertyRegistry = "REPEAT"
	// TriggerProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3
	TriggerProp PropertyRegistry = "TRIGGER"
	// DateTimeCreatedProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.1
	DateTimeCreatedProp PropertyRegistry = "CREATED"
	// DateTimeStampProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2
	DateTimeStampProp PropertyRegistry = "DTSTAMP"
	// LastModifiedProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3
	LastModifiedProp PropertyRegistry = "LAST-MODIFIED"
	// SequenceProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4
	SequenceProp PropertyRegistry = "SEQUENCE"
	// RequestStatusProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.3
	RequestStatusProp PropertyRegistry = "REQUEST-STATUS"
	// NonStandardProp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.2
	NonStandardProp PropertyRegistry = "NON-STANDARD-PROPERTY"

	// BeginProp is used to start a new component definition
	BeginProp PropertyRegistry = "BEGIN"
	// EndProp is used to end a component definition
	EndProp PropertyRegistry = "END"
)
