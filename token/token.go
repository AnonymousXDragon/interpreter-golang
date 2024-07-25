package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers
	IDENT = "IDENTIFIER"
	INT   = "INT"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	DIVIDE   = "/"
	ASTERISK = "*"

	// delimiters
	SEMICOLUMN = ";"
	COMMA      = ","

	LPAREN = "("
	RPAREN = ")"

	RBRACE = "{"
	LBRACE = "}"

	// conditional operators
	NOT         = "!"
	GRETER_THAN = ">"
	LESS_THAN   = "<"

	// Two Character Token
	EQUAL     = "=="
	NOT_EQUAL = "!="

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
	TRUE     = "true"
	FALSE    = "false"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookUp(char string) TokenType {
	if t, ok := keywords[char]; ok {
		return t
	}
	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(Ttype TokenType, val byte) Token {
	return Token{
		Type:    Ttype,
		Literal: string(val),
	}
}
