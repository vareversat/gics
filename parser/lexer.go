package parser

type parameterParser struct {
	ParameterName  string
	ParameterValue string
}

type PropertyParser struct {
	PropertyName     string
	PropertyValue    string
	ParameterParsers []parameterParser
}

func NewPropertyParser() *PropertyParser {
	return &PropertyParser{
		PropertyName:     "",
		PropertyValue:    "",
		ParameterParsers: make([]parameterParser, 0),
	}
}

func (p *PropertyParser) SetPropertyName(name string) {
	p.PropertyName = name
}

func (p *PropertyParser) SetPropertyValue(value string) {
	p.PropertyValue = value
}

func (p *PropertyParser) AddParam(paramName string) {
	p.ParameterParsers = append(
		p.ParameterParsers,
		parameterParser{ParameterName: paramName, ParameterValue: ""},
	)
}

func (p *PropertyParser) SetValueOfLastParam(paramValue string) {
	p.ParameterParsers[len(p.ParameterParsers)-1].ParameterValue = paramValue
}

type Lexer struct {
	input         string
	position      int  // current position in input (points to current char)
	readPosition  int  // current reading position in input (after current char)
	ch            byte // current char under examination
	previousCh    byte // previous char examined
	property      *PropertyParser
	isRelaxedMode bool
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, isRelaxedMode: false}
	l.property = NewPropertyParser()
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token
	switch l.ch {
	case '=':
		tok = newToken(Equal, l.ch)
	case ':':
		tok = newToken(Colon, l.ch)
	case ';':
		tok = newToken(Semicolon, l.ch)
	case '"':
		l.isRelaxedMode = !l.isRelaxedMode
		tok = newToken(DoubleQuote, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		// We started to reader the property name
		if l.position == 0 {
			tok.Literal = l.readIdentifier()
			tok.Type = PropertyNameToken
			l.property.SetPropertyName(tok.Literal)
			return tok
		}
		// We started reading property value
		if l.previousCh == ':' {
			// When reading prop value, you have to be as relaxed as possible to keep all special characters
			l.isRelaxedMode = true
			tok.Literal = l.readIdentifier()
			l.isRelaxedMode = false
			tok.Type = PropertyValueToken
			l.property.SetPropertyValue(tok.Literal)
			return tok
		}
		// We started to reader a parameter name
		if l.previousCh == ';' {
			tok.Literal = l.readIdentifier()
			tok.Type = ParameterNameToken
			l.property.AddParam(tok.Literal)
			return tok
		}
		// We started to reader a parameter value
		if l.previousCh == '=' {
			tok.Literal = l.readIdentifier()
			tok.Type = ParameterValueToken
			l.property.SetValueOfLastParam(tok.Literal)
			return tok
		}
		// We started to reader a double-quoted parameter name
		if l.previousCh == '"' {
			tok.Literal = l.readIdentifier()
			tok.Type = ParameterValueToken
			l.property.SetValueOfLastParam(tok.Literal)
			return tok
		}
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
		if l.readPosition > 0 {
			l.previousCh = l.input[l.readPosition-1]
		}
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	// We run into a '"' or we are reading a property value
	if l.isRelaxedMode {
		for isLetterRelaxed(l.ch) {
			l.readChar()
		}
	} else {
		for isLetter(l.ch) {
			l.readChar()
		}
	}
	return l.input[position:l.position]
}

func isLetterRelaxed(ch byte) bool {
	return ch != '"' && ch != 0
}

func isLetter(ch byte) bool {
	return ch != ':' && ch != ';' && ch != '=' && ch != 0
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
