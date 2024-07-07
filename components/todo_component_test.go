package components

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/properties"
)

func TestNewToDoCalendarComponent(t *testing.T) {
	component := NewToDoCalendarComponent(
		properties.NewUidProperty("UID"),
		properties.NewDateTimeStampProperty(time.Now()),
		[]AlarmCalendarComponent{},
	)

	assert.NotNil(t, component)
}

func TestToDoCalendarComponent_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	uid := properties.NewUidProperty("My ToDo")
	dateTimestamp := properties.NewDateTimeStampProperty(myTime)

	component := NewToDoCalendarComponent(uid, dateTimestamp, []AlarmCalendarComponent{})

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VTODO\r\nUID:My ToDo\r\nDTSTAMP:%s\r\nEND:VTODO\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}
