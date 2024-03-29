package pkg

import (
	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type Calendars []calendarComponent

type calendarComponent struct {
	Begin properties.ComponentBeginProperty
	End   properties.ComponentEndProperty
	UID   properties.UidProperty

	DateTimeStamp      properties.DateTimeStampProperty
	DateTimeStart      properties.DateTimeStartProperty
	Class              properties.ClassificationProperty
	DateTimeCreated    properties.DateTimeCreatedProperty
	Description        properties.DescriptionProperty
	GeographicPosition properties.GeographicPositionProperty
	LastModified       properties.LastModifiedProperty
	Location           properties.LocationProperty
	Organize           properties.OrganizerProperty
	Priority           properties.PriorityProperty
	Sequence           properties.SequenceProperty
	Status             properties.StatusProperty
	Summary            properties.SummaryProperty
	TimeTransparency   properties.TimeTransparencyProperty
	Url                properties.UrlProperty
	RecurrenceId       properties.RecurrenceIdProperty

	// RRULE

	DateTimeEnd properties.DateTimeEndProperty
	Duration    properties.DurationProperty

	Attachment        properties.AttachmentProperty
	Attendee          properties.AttendeeProperty
	Categories        properties.CategoriesProperty
	Comment           properties.CommentProperty
	Contact           properties.ContactProperty
	ExceptionDateTime properties.ExceptionDateTimeProperty
	RequestStatus     properties.RequestStatusProperty
	RelatedTo         properties.RelatedToProperty
	Resources         properties.ResourcesProperty
	// RDATE
}

type EventComponent interface {
}

func NewEventComponent() EventComponent {
	return &calendarComponent{
		Begin: properties.NewComponentBeginProperty(registries.VEVENT),
		UID:   properties.NewUidProperty("my_first_event"),
		End:   properties.NewComponentEndProperty(registries.VEVENT),
	}
}
