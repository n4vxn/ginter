package lexer

import (
	"fmt"
	"go/token"
)

type TokenType string

const PLUS = "+"

type Lexer struct {
	input        string
	position     int  // curr position in input (points to curr char)
	readPosition int  // curr reading position in input (after curr char)
	ch           byte // curr char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS_SIGN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar
	return tok
}

func newToken(tokenType token.Token, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
