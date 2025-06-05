package main

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() bool {
	return p.parseObject() && p.expect(EOF)
}

func (p *Parser) parseValue() bool {
	switch p.peek().Type {
	case STRING, NUMBER, TRUE, FALSE, NULL:
		p.next()
		return true
	case LBRACE:
		return p.parseObject()
	case LBRACKET:
		return p.parseArray()
	default:
		return false
	}
}

func (p *Parser) parseObject() bool {
	if !p.expect(LBRACE) {
		return false
	}

	if p.peek().Type == RBRACE {
		p.next()
		return true
	}

	for {
		if !p.expect(STRING) || !p.expect(COLON) {
			return false
		}
		if !p.parseValue() {
			return false
		}
		if p.peek().Type == COMMA {
			p.next()
			continue
		} else if p.peek().Type == RBRACE {
			p.next()
			break
		} else {
			return false
		}
	}
	return true
}

func (p *Parser) parseArray() bool {
	if !p.expect(LBRACKET) {
		return false
	}
	if p.peek().Type == RBRACKET {
		p.next()
		return true
	}

	for {
		if !p.parseValue() {
			return false
		}
		if p.peek().Type == COMMA {
			p.next()
			continue
		} else if p.peek().Type == RBRACKET {
			p.next()
			break
		} else {
			return false
		}
	}

	return true
}

func (p *Parser) peek() Token {
	if p.pos >= len(p.tokens) {
		return Token{Type: EOF}
	}
	return p.tokens[p.pos]
}

func (p *Parser) next() {
	p.pos++
}

func (p *Parser) expect(t TokenType) bool {
	if p.peek().Type == t {
		p.next()
		return true
	}
	return false
}
