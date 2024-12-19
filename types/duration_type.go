package types

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/vareversat/gics/registries"
)

const (
	yearToNs   int = 31557600 * 1e9
	monthToNs  int = 2628003 * 1e9
	weekToNs   int = 604800 * 1e9
	dayToNs    int = 86400 * 1e9
	hourToNs   int = 3600 * 1e9
	minuteToNs int = 60 * 1e9
	secondToNs int = 1e9
)

type DurationType interface {
	ValueType
	GetValue() time.Duration
}

type durationType struct {
	typeName  registries.ValueTypeRegistry
	typeValue time.Duration
}

// parseDurationToString return the correct representation from a duration (ex: PT5H30M)
func parseDurationToString(duration time.Duration) string {
	finalString := "P"
	// Get the absolute value (in case of negative duration)
	toAbsNs := math.Abs(float64(duration.Nanoseconds()))
	years := int(toAbsNs) / yearToNs
	// 1 mont = 2628003 secs
	months := (int(toAbsNs) % yearToNs) / monthToNs
	// 1 week = 604800 secs
	weeks := ((int(toAbsNs) % yearToNs) % monthToNs) / weekToNs
	// 1 day = 86400 secs
	days := (((int(toAbsNs) % yearToNs) % monthToNs) % weekToNs) / dayToNs
	// 1 hour = 3600 secs
	hours := ((((int(toAbsNs) % yearToNs) % monthToNs) % weekToNs) % dayToNs) / hourToNs
	// 1 minute = 60 secs
	minutes := (((((int(toAbsNs) % yearToNs) % monthToNs) % weekToNs) % dayToNs) % hourToNs) / minuteToNs
	// seconds = the remainder
	seconds := ((((((int(toAbsNs) % yearToNs) % monthToNs) % weekToNs) % dayToNs) % hourToNs) % minuteToNs) / secondToNs

	// Test if the duration is negative or not
	if math.Signbit(duration.Seconds()) {
		finalString = "-" + finalString
	}
	if years != 0 {
		finalString += fmt.Sprintf("%dY", years)
	}
	if months != 0 {
		finalString += fmt.Sprintf("%dM", months)
	}
	if weeks != 0 {
		finalString += fmt.Sprintf("%dW", weeks)
	}
	if days != 0 {
		finalString += fmt.Sprintf("%dD", days)
	}
	if hours != 0 || minutes != 0 || seconds != 0 {
		finalString += "T"
	}
	if hours != 0 {
		finalString += fmt.Sprintf("%dH", hours)
	}
	if minutes != 0 {
		finalString += fmt.Sprintf("%dM", minutes)
	}
	if seconds != 0 {
		finalString += fmt.Sprintf("%dS", seconds)
	}

	return finalString
}

func (t *durationType) GetStringValue() string {
	return parseDurationToString(t.typeValue)
}

func (t *durationType) GetTypeName() string {
	return string(t.typeName)
}

// GetValue get the [time.Duration] typed value
func (t *durationType) GetValue() time.Duration {
	return t.typeValue
}

// parseStringToDuration With a given text value (ex: P2W6DT6H0M0S) return a [time.Duration] value
func parseStringToDuration(value string) time.Duration {
	var finalDuration time.Duration
	currentQuantity := ""
	readingTime := false
	reader := bufio.NewReader(strings.NewReader(value))
	for {
		if r, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			str := string(r)
			// Try to convert the current rune to a int
			_, err := strconv.Atoi(str)
			if err == nil {
				// If so, that mean we hit a quantity. Concat to the existing string formatted number
				currentQuantity += str
			}
			// Else, test if we hit a special character that indicate the end of a quantity
			if str == "T" {
				// We are entering in time reading mode
				readingTime = true
			}
			if str == "Y" {
				value, _ := strconv.Atoi(currentQuantity)
				finalDuration += time.Duration(value * yearToNs)
				currentQuantity = ""
			}
			if str == "M" && !readingTime {
				value, _ := strconv.Atoi(currentQuantity)
				finalDuration += time.Duration(value * monthToNs)
				currentQuantity = ""
			}
			if str == "W" {
				value, _ := strconv.Atoi(currentQuantity)
				finalDuration += time.Duration(value * weekToNs)
				currentQuantity = ""
			}
			if str == "D" {
				value, _ := strconv.Atoi(currentQuantity)
				finalDuration += time.Duration(value * dayToNs)
				currentQuantity = ""
			}
			if str == "H" {
				value, _ := strconv.Atoi(currentQuantity)
				finalDuration += time.Duration(value * hourToNs)
				currentQuantity = ""
			}
			if str == "M" && readingTime {
				value, _ := strconv.Atoi(currentQuantity)
				finalDuration += time.Duration(value * minuteToNs)
				currentQuantity = ""
			}
			if str == "S" {
				value, _ := strconv.Atoi(currentQuantity)
				finalDuration += time.Duration(value * secondToNs)
				currentQuantity = ""
			}
		}
	}
	return finalDuration
}

// NewDurationValue create a new [registries.Duration] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
func NewDurationValue(value time.Duration) DurationType {
	return &durationType{
		typeName:  registries.Duration,
		typeValue: value,
	}
}

// NewDurationValue create a new [registries.Duration] type value from a string. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
func NewDurationValueFromString(value string) DurationType {
	return &durationType{
		typeName:  registries.Duration,
		typeValue: parseStringToDuration(value),
	}
}
