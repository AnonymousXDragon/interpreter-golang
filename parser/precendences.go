package parser

import (
	"brLang/token"
	"go/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // < or >
	SUM         // +
	PRODUCT     //*
	PREFIX      // - or !
	CALL        // function call
)

var precendence = map[token.TokenType]int{
	token.EQUAL:       EQUALS,
	token.NOT_EQUAL:   EQUALS,
	token.LESS_THAN:   LESSGREATER,
	token.GRETER_THAN: LESSGREATER,
	token.PLUS:        SUM,
	token.MINUS:       SUM,
	token.DIVIDE:      PRODUCT,
	token.ASTERISK:    PRODUCT,
}

func (p *Parser) PeekPrecedence() int {
	if t, ok := precendence[p.peekToken.Type]; ok {
		return t
	}
	return LOWEST
}

func (p *Parser) CurrentPrecedence() int {
	if t, ok := precendence[p.currentToken.Type]; ok {
		return t
	}
	return LOWEST
}
