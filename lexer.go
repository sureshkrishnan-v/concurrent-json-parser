package main

import (
	"strings"
)

type TokenType string

const (
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	ILLEGAL  TokenType = "ILLEGAL"
	EOF      TokenType = "EOF"
	STRING   TokenType = "STRING"
	COLON    TokenType = ":"
	COMMA    TokenType = ","
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	NUMBER   TokenType = "NUMBER"
	NULL     TokenType = "NULL"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: strings.TrimSpace(input)}
}

func (l *Lexer) Lex() []Token {
	var tokens []Token

	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		switch ch {
		case '{':
			tokens = append(tokens, Token{
				Type:    LBRACE,
				Literal: "{",
			})
			l.pos++
		case '}':
			tokens = append(tokens,
				Token{
					Type:    RBRACE,
					Literal: "}"},
			)
			l.pos++
		case ' ', '\t', '\n', '\r':
			l.pos++
			//ignore
		case ':':
			tokens = append(tokens, Token{
				Type:    COLON,
				Literal: ":",
			})
			l.pos++
		case ',':
			tokens = append(tokens, Token{
				Type:    COMMA,
				Literal: ",",
			})
			l.pos++
		case '[':
			tokens = append(tokens, Token{
				Type: LBRACKET, Literal: "[",
			})
			l.pos++
		case ']':
			tokens = append(tokens, Token{
				Type: RBRACKET, Literal: "]",
			})
			l.pos++

		case '"':
			str := l.readString()
			tokens = append(tokens, Token{
				Type:    STRING,
				Literal: str,
			})

		default:
			if isDigit(ch) {
				tokens = append(tokens, Token{
					Type: NUMBER, Literal: l.readNumber(),
				})
			} else if isAlpha(ch) {
				ident := l.readIdentifier()
				switch ident {
				case "null":
					tokens = append(tokens, Token{
						Type: NULL, Literal: "null",
					})
				case "true":
					tokens = append(tokens, Token{
						Type: TRUE, Literal: "true",
					})
				case "false":
					tokens = append(tokens, Token{
						Type: FALSE, Literal: "false",
					})
				default:
					tokens = append(tokens, Token{
						Type: ILLEGAL, Literal: ident,
					})

				}
			} else {
				tokens = append(tokens, Token{
					Type:    ILLEGAL,
					Literal: string(ch),
				})
				l.pos++
			}
		}
	}
	tokens = append(tokens, Token{
		Type: EOF,
	})
	return tokens
}

func (l *Lexer) readString() string {
	start := l.pos + 1
	for i := start; i < len(l.input); i++ {
		if l.input[i] == '"' {
			str := l.input[start:i]
			l.pos = i + 1
			return str
		}
	}
	l.pos = len(l.input)
	return ""
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
		l.pos++
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readIdentifier() string {
	start := l.pos
	for l.pos < len(l.input) && isAlpha(l.input[l.pos]) {
		l.pos++
	}
	return l.input[start:l.pos]
}
