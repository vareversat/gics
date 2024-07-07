package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	escapeCharacter     = "\r\n"
	foldedLineBeginning = " "
	foldingIndex        = 75
	foldingCharacter    = escapeCharacter + foldedLineBeginning
)

// unfold will scan all the file and replace "\r\n " by "\r\n"
// Take a os.File and populate a bytes.Buffer pointer
func unfold(foldedFile *os.File, unfoldedData *bytes.Buffer) {
	scanner := bufio.NewScanner(foldedFile)
	// Iterate through all the folded file
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, foldedLineBeginning) {
			// Remove the last 2 characters of our buffer (\r\n) if we encounter a " " in a new line
			previousLineWithoutEscapeCharacter := bytes.Clone(
				unfoldedData.Bytes()[:len(unfoldedData.Bytes())-2],
			)
			unfoldedData.Reset()
			unfoldedData.Write(previousLineWithoutEscapeCharacter)
			// Insert this new line minus the " " character and then add the line break character (\r\n)
			unfoldedData.WriteString(fmt.Sprintf("%s%s", line[1:], escapeCharacter))
		} else {
			// Append the line to our unfolded data pointer
			unfoldedData.WriteString(fmt.Sprintf("%s%s", line, escapeCharacter))
		}
	}
}

func parse(unfoldedData bytes.Buffer) {
	unfoldedDataReader := bytes.NewReader(unfoldedData.Bytes())
	scanner := bufio.NewScanner(unfoldedDataReader)
	// Iterate through all the unfolded file
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lexer := NewLexer(line)
		for tok := lexer.NextToken(); tok.Type != EOF; tok = lexer.NextToken() {
		}
		parsedProperty, _ := ParseProperty(
			lexer.property.PropertyName,
			lexer.property.PropertyValue,
		)
		if parsedProperty != nil {
			parsedProperty.ToICalendarPropFormat(os.Stdout)
		}
		lineNumber++
	}
}
