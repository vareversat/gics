package main

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/vareversat/gics/pkg"
	"github.com/vareversat/gics/pkg/parser"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types/recurrence_rule"

	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/types"

	"github.com/vareversat/gics/pkg/components"
)

func main() {

	// VEVENT
	url1, _ := url.Parse("CID:part3.msg.970415T083000@example.com")
	url2, _ := url.Parse("mailto:test@example.com")
	url3, _ := url.Parse("https://example.com")
	url4, _ := url.Parse("mailto:someone@company.com")

	vevent := components.NewEventCalendarComponent(
		properties.NewDescriptionProperty(
			"This is a description sdklhjfkdjshfjkqsdhfkjqshddkfjhkuliueuifhj<ehfkj<dshfjkshfuhzuhfkkjshkjdhfjksqhfdk",
			parameters.NewLanguageParam("EN"),
			parameters.NewAlternateTextRepresentationParam(url1),
		),
		properties.NewGeographicPositionProperty(37.386013, -122.082932),
		properties.NewLastModifiedProperty(time.Now()),
		properties.NewLocationProperty("Paris",
			parameters.NewLanguageParam("EN"),
			parameters.NewAlternateTextRepresentationParam(url1)),
		properties.NewOrganizerProperty(url2,
			parameters.NewLanguageParam("EN"),
			parameters.NewCommonNameParam("Common Name"),
			parameters.NewDirectoryEntryParam(url2),
			parameters.NewSentByParam(url2)),
		properties.NewPriorityProperty(1),
		properties.NewSequenceProperty(0),
		properties.NewStatusProperty(types.VEVENT_CONFIRMED),
		properties.NewSummaryProperty("This a summary",
			parameters.NewLanguageParam("EN"),
			parameters.NewAlternateTextRepresentationParam(url1)),
		properties.NewTimeTransparencyProperty(types.OPAQUE),
		properties.NewUrlProperty(url3),
		properties.NewRecurrenceIdProperty(time.Now(), types.WithLocalTime,
			parameters.NewTimeZoneIdentifierParam("Europe/Paris"),
			parameters.NewValueParam(registries.DATE)),
		properties.NewDateTimeEndProperty(time.Now(), types.WithLocalTime,
			parameters.NewTimeZoneIdentifierParam("Europe/Paris"),
			parameters.NewValueParam(registries.DATE)),
		properties.NewRecurrenceRuleProperty(
			recurrence_rule.NewFrequencyPart(recurrence_rule.YEARLY),
			recurrence_rule.NewIntervalPart(2)),
		properties.NewDurationProperty("PT1H0M0S"),
		properties.NewAttachmentProperty(
			"VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=",
			parameters.NewFormatTypeParam("text/plain"),
			parameters.NewInlineEncodingParam(parameters.BASE64),
		),
		properties.NewAttendeeProperty(url4,
			parameters.NewCalendarUserTypeParam(parameters.INDIVIDUAL),
			parameters.NewMemberParam(url4),
			parameters.NewParticipationRoleParam(parameters.CHAIR),
			parameters.NewParticipationStatusParam(parameters.NEEDACTION),
			parameters.NewRSVPParam(true),
			parameters.NewDelegateesParam(url4),
			parameters.NewDelegatorsParam(url4),
			parameters.NewSentByParam(url4),
			parameters.NewCommonNameParam("Common Name"),
			parameters.NewDirectoryEntryParam(url4),
			parameters.NewLanguageParam("EN")),
		properties.NewCategoryProperty([]string{"APPOINTMENT", "EDUCATION"},
			parameters.NewLanguageParam("EN")),
		properties.NewCommentProperty("This is a comment",
			parameters.NewLanguageParam("EN"),
			parameters.NewAlternateTextRepresentationParam(url1)),
		properties.NewContactProperty("This is a contact",
			parameters.NewLanguageParam("EN"),
			parameters.NewAlternateTextRepresentationParam(url1)),
		properties.NewExceptionDateTimeProperty(
			[]time.Time{time.Now(), time.Now()},
			types.WithUtcTime,
			parameters.NewTimeZoneIdentifierParam("Europe/Paris"),
			parameters.NewValueParam(registries.DATETIME),
		),
		properties.NewRequestStatusProperty("2.0", "Success", "",
			parameters.NewLanguageParam("EN")),
		properties.NewRelatedToProperty("Relates to another calendar",
			parameters.NewRelationshipParam(parameters.CHILD)),
		properties.NewResourcesProperty([]string{"CAR", "TRUCK"},
			parameters.NewLanguageParam("EN"),
			parameters.NewAlternateTextRepresentationParam(url1)),
		properties.NewRecurrenceDateTimesProperty([]time.Time{time.Now(), time.Now()},
			[]time.Time{time.Now(), time.Now()}, types.WithUtcTime,
			parameters.NewTimeZoneIdentifierParam("Europe/Paris"),
			parameters.NewValueParam(registries.PERIOD)),
	)

	component := make(components.Components, 0)

	component = append(component, vevent)
	calendar, err := pkg.NewCalendar(component, "this_project", "xyz", "2.0")
	event, err := os.OpenFile("event.ics", os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier de sortie:", err)
		return
	}
	calendar.ToICalendarFormat(event)
	event.Close()
	parser.Parser()

}
