package components

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/pkg/properties"
)

func TestNewJournalCalendarComponent(t *testing.T) {
	component := NewJournalCalendarComponent(
		properties.NewUidProperty("UID"),
		properties.NewDateTimeStampProperty(time.Now()),
	)

	assert.NotNil(t, component)
}

func TestJournalCalendarComponent_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	uid := properties.NewUidProperty("My Journal")
	dateTimestamp := properties.NewDateTimeStampProperty(myTime)

	component := NewJournalCalendarComponent(uid, dateTimestamp)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VJOURNAL\r\nUID:My Journal\r\nDTSTAMP:%s\r\nEND:VJOURNAL\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}
