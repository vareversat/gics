package components

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/types"
)

func TestNewEventCalendarComponent(t *testing.T) {
	action := properties.NewActionProperty(types.AUDIO)
	trigger := properties.NewTriggerProperty(time.Now(), "PT5M", types.WithUtcTime)
	component := NewEventCalendarComponent(action, trigger)

	assert.NotNil(t, component)
	assert.Equal(t, action.GetValue(), types.AUDIO)
}

func TestEventCalendarComponent_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	uid := properties.NewUidProperty("My Event")
	dateTimestamp := properties.NewDateTimeStampProperty(myTime, types.WithUtcTime)
	description := properties.NewDescriptionProperty("Test Event")
	summary := properties.NewSummaryProperty("Event Summary")

	component := NewEventCalendarComponent(uid, dateTimestamp, description, summary)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VEVENT\r\nUID:My Event\r\nDTSTAMP:%s\r\nDESCRIPTION:Test Event\r\nSUMMARY:Event Summary\r\nEND:VEVENT\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}
