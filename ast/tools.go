package ast

import (
	"brLang/token"
	"strings"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// let statement
func (l *LetStatement) StatementNode() {}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}
func (l *LetStatement) String() string {
	var out strings.Builder

	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

// expression
func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

// return statement
func (rs *ReturnStatement) StatementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out strings.Builder

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

// expression statement
func (es *ExpressionStatement) StatementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// integer literal
func (in *IntegerLiteral) ExpressionNode() {}
func (in *IntegerLiteral) TokenLiteral() string {
	return in.Token.Literal
}
func (in *IntegerLiteral) String() string {
	return in.Token.Literal
}

// prefix expression
func (pf *PrefixExpression) ExpressionNode() {}
func (pf *PrefixExpression) TokenLiteral() string {
	return pf.Token.Literal
}
func (pf *PrefixExpression) String() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(pf.Operator)
	out.WriteString(pf.Right.String())
	out.WriteString(")")

	return out.String()
}

// infix expression
func (ne *InfixExpression) ExpressionNode() {}
func (ne *InfixExpression) TokenLiteral() string {
	return ne.Token.Literal
}
func (ne *InfixExpression) String() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(ne.Left.String())
	out.WriteString(" " + ne.Operator + " ")
	out.WriteString(ne.Right.String())
	out.WriteString(")")

	return out.String()
}
