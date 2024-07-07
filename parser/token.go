package parser

type TokenType string

const (
	PropName     TokenType = "PropName"
	PropValue    TokenType = "PropValue"
	ParamName    TokenType = "ParamName"
	ParamValue   TokenType = "ParamValue"
	ParamStatus  TokenType = "ParamValue"
	Equal        TokenType = "EQUAL"
	Semicolon    TokenType = "SEMICOLON"
	Colon        TokenType = "COLON"
	Comma        TokenType = ","
	ForwardSlash TokenType = "/"
	DoubleQuote  TokenType = "DoubleQuote"
	EOF          TokenType = "EOF"
	ILLEGAL      TokenType = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
}
