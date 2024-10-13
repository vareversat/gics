package parser

type TokenType string

const (
	PropertyNameToken   TokenType = "PropName"
	PropertyValueToken  TokenType = "PropValue"
	ParameterNameToken  TokenType = "ParamName"
	ParameterValueToken TokenType = "ParamValue"
	ParameterStatus     TokenType = "ParamValue"
	Equal               TokenType = "EQUAL"
	Semicolon           TokenType = "SEMICOLON"
	Colon               TokenType = "COLON"
	Comma               TokenType = ","
	ForwardSlash        TokenType = "/"
	DoubleQuote         TokenType = "DoubleQuote"
	EOF                 TokenType = "EOF"
	ILLEGAL             TokenType = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
}
