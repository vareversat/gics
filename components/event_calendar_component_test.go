package components

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/vareversat/gics/registries"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/properties"
)

func TestNewEventCalendarComponent(t *testing.T) {
	component := NewEventCalendarComponent(
		properties.NewUidProperty("UID"),
		properties.NewDateTimeStampProperty(time.Now()),
	)

	assert.NotNil(t, component)
}

func TestEventCalendarComponent_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	uid := properties.NewUidProperty("My Event")
	dateTimestamp := properties.NewDateTimeStampProperty(myTime)
	description := properties.NewDescriptionProperty("Test Event")
	summary := properties.NewSummaryProperty("Event SummaryProp")

	component := NewEventCalendarComponent(
		uid,
		dateTimestamp,
		description,
		summary,
	)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VEVENT\r\nUID:My Event\r\nDTSTAMP:%s\r\nDESCRIPTION:Test Event\r\nSUMMARY:Event SummaryProp\r\nEND:VEVENT\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}

func TestEventCalendarComponent_WithAlarm(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	action := properties.NewActionProperty(registries.Audio)
	trigger, _ := properties.NewTriggerFromDurationProperty(-3600 * 1e9)
	alarmDescription := properties.NewDescriptionProperty("Test Alarm")
	alarmSummary := properties.NewSummaryProperty("Alarm SummaryProp")

	alarm := NewAlarmCalendarComponent(action, trigger, alarmDescription, alarmSummary)

	uid := properties.NewUidProperty("My Event")
	dateTimestamp := properties.NewDateTimeStampProperty(myTime)
	description := properties.NewDescriptionProperty("Test Event")
	summary := properties.NewSummaryProperty("Event SummaryProp")

	component := NewEventCalendarComponent(
		uid,
		dateTimestamp,
		description,
		summary,
	)

	component.AddAlarmComponent(alarm)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VEVENT\r\nUID:My Event\r\nDTSTAMP:%s\r\nDESCRIPTION:Test Event\r\nSUMMARY:Event SummaryProp\r\nBEGIN:VALARM\r\nACTION:AUDIO\r\nTRIGGER:-PT1H\r\nDESCRIPTION:Test Alarm\r\nSUMMARY:Alarm SummaryProp\r\nEND:VALARM\r\nEND:VEVENT\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}

func TestEventCalendarComponent_GetProperty(t *testing.T) {
	myTime := time.Now()

	uid := properties.NewUidProperty("My Event")
	dateTimestamp := properties.NewDateTimeStampProperty(myTime)
	description := properties.NewDescriptionProperty("Test Event")
	summary := properties.NewSummaryProperty("Event SummaryProp")

	component := NewEventCalendarComponent(
		uid,
		dateTimestamp,
		description,
		summary,
	)

	prop := component.GetProperty(registries.SummaryProp)
	assert.Equal(t, "Event SummaryProp", prop.GetValue())
}
