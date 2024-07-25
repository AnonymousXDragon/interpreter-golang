package lexer

import (
	tokens "brLang/token"
)

type Lexer struct {
	Input           string
	CurrentPosition int
	NextPosition    int
	Char            byte
}

func New(input string) *Lexer {
	l := &Lexer{
		Input: input,
	}
	l.Read()
	return l
}

func (l *Lexer) Read() {
	if l.NextPosition >= len(l.Input) {
		l.Char = 0
	} else {
		l.Char = l.Input[l.NextPosition]
	}
	l.CurrentPosition = l.NextPosition
	l.NextPosition += 1
}

func (l *Lexer) PeekRead() byte {
	if l.NextPosition > len(l.Input) {
		return 0
	} else {
		return l.Input[l.NextPosition]
	}
}

func (l *Lexer) NextToken() tokens.Token {
	var token tokens.Token

	l.EatWhitespace()

	switch l.Char {
	case '=':
		if l.PeekRead() == '=' {
			ch := l.Char
			l.Read()
			token = tokens.Token{
				Type:    tokens.EQUAL,
				Literal: string(ch) + string(l.Char),
			}
		} else {
			token = tokens.NewToken(tokens.ASSIGN, l.Char)
		}

	case ';':
		token = tokens.NewToken(tokens.SEMICOLUMN, l.Char)
	case '(':
		token = tokens.NewToken(tokens.LPAREN, l.Char)
	case ')':
		token = tokens.NewToken(tokens.RPAREN, l.Char)
	case ',':
		token = tokens.NewToken(tokens.COMMA, l.Char)
	case '+':
		token = tokens.NewToken(tokens.PLUS, l.Char)
	case '{':
		token = tokens.NewToken(tokens.RBRACE, l.Char)
	case '}':
		token = tokens.NewToken(tokens.LBRACE, l.Char)
	case '-':
		token = tokens.NewToken(tokens.MINUS, l.Char)
	case '/':
		token = tokens.NewToken(tokens.DIVIDE, l.Char)
	case '*':
		token = tokens.NewToken(tokens.ASTERISK, l.Char)
	case '!':
		if l.PeekRead() == '=' {
			ch := l.Char
			l.Read()
			token = tokens.Token{
				Type:    tokens.NOT_EQUAL,
				Literal: string(l.Char) + string(ch),
			}
		} else {
			token = tokens.NewToken(tokens.NOT, l.Char)
		}
	case 0:
		token.Literal = ""
		token.Type = tokens.EOF
	default:
		if l.IsLetter(l.Char) {
			token.Literal = l.ReadIdentifier()
			token.Type = tokens.LookUp(token.Literal)
			return token
		} else if l.IsDigit(l.Char) {
			token.Type = tokens.INT
			token.Literal = l.ReadNumber()
			return token
		} else {
			token = tokens.NewToken(tokens.ILLEGAL, l.Char)
		}
	}

	l.Read()
	return token
}

func (l *Lexer) IsLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'B'
}

func (l *Lexer) ReadIdentifier() string {
	position := l.CurrentPosition
	for l.IsLetter(l.Char) {
		l.Read()
	}
	return l.Input[position:l.CurrentPosition]
}

func (l *Lexer) ReadNumber() string {
	position := l.CurrentPosition
	for l.IsDigit(l.Char) {
		l.Read()
	}

	return l.Input[position:l.CurrentPosition]
}

func (l *Lexer) IsDigit(char byte) bool {
	return '0' <= l.Char && l.Char <= '9'
}

func (l *Lexer) EatWhitespace() {
	for l.Char == ' ' || l.Char == '\n' || l.Char == '\t' || l.Char == '\r' {
		l.Read()
	}
}
