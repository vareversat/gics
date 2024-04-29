package components

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/pkg/properties"
)

func TestNewEventCalendarComponent(t *testing.T) {
	component := NewEventCalendarComponent(
		properties.NewUidProperty("UID"),
		properties.NewDateTimeStampProperty(time.Now()),
		[]AlarmCalendarComponent{},
	)

	assert.NotNil(t, component)
}

func TestEventCalendarComponent_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	uid := properties.NewUidProperty("My Event")
	dateTimestamp := properties.NewDateTimeStampProperty(myTime)
	description := properties.NewDescriptionProperty("Test Event")
	summary := properties.NewSummaryProperty("Event Summary")

	component := NewEventCalendarComponent(
		uid,
		dateTimestamp,
		[]AlarmCalendarComponent{},
		description,
		summary,
	)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VEVENT\r\nUID:My Event\r\nDTSTAMP:%s\r\nDESCRIPTION:Test Event\r\nSUMMARY:Event Summary\r\nEND:VEVENT\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}
