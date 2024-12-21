package parser

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestEventCalendarComponentParsing_1(t *testing.T) {
	input, err := os.OpenFile("./test_data/event_1.ics", os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer input.Close()
	unfoldedData := bytes.Buffer{}

	unfold(input, &unfoldedData)
	parse(unfoldedData)
}

func TestEventCalendarComponentParsing_2(t *testing.T) {
	input, err := os.OpenFile("./test_data/event_2.ics", os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer input.Close()
	unfoldedData := bytes.Buffer{}

	unfold(input, &unfoldedData)
	parse(unfoldedData)
}

func TestEventCalendarComponentParsing_3(t *testing.T) {
	input, err := os.OpenFile("./test_data/event_3.ics", os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer input.Close()
	unfoldedData := bytes.Buffer{}

	unfold(input, &unfoldedData)
	parse(unfoldedData)
}
