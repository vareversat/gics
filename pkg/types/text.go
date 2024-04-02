package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.11

type TextValue struct {
	S string
}

func NewTextValue(value string) TextValue {
	return TextValue{
		S: value,
	}
}

func NewTextValues(values []string) []TextValue {
	textValues := make([]TextValue, 0)

	for i := 0; i < len(values); i++ {
		textValues = append(textValues, TextValue{
			S: values[i],
		})
	}
	return textValues
}
