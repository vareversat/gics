package parser

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestNewAlarmCalendarComponent(t *testing.T) {
	input, err := os.OpenFile("./test_data/event.ics", os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer input.Close()
	unfoldedData := bytes.Buffer{}

	unfold(input, &unfoldedData)
	parse(unfoldedData)
}
