package template

import (
	"fmt"
)

type token struct {
	typ tokenType
	val string
}

type parser struct {
	tokens chan token // channel of parsed token
	items  chan item  // chan to send tokens for client
	buf    *item      // have a buffer of 1 item for parser
	lex    *lexer     // input lexer
}

type tokenType int

const (
	tokenError tokenType = iota
	tokenEnd
	tokenIdentifier
	tokenText
	tokenLeftDelim
	tokenRightDelim
)

type parseStateFn func(*parser) parseStateFn

func parse(input, leftDelim, rightDelim string) *parser {
	p := &parser{
		tokens: make(chan token),
		lex:    lex(input, leftDelim, rightDelim),
	}
	go p.run()
	return p
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state.
func (p *parser) errorf(format string, args ...interface{}) parseStateFn {
	p.tokens <- token{tokenError, fmt.Sprintf(format, args...)}
	return nil
}

// emit passes a token back to the client.
func (p *parser) emit(t token) {
	p.tokens <- t
}

// nextItem returns the next item when it becomes available
func (p *parser) nextItem() token {
	return <-p.tokens
}

// peek returns but does not consume the next token.
func (p *parser) peek() item {
	if p.buf == nil {
		i := p.lex.nextItem()
		p.buf = &i
	}
	return *p.buf
}

func (p *parser) getBufOrNext() item {
	if p.buf != nil {
		item := *p.buf
		p.buf = nil
		return item
	}
	return p.lex.nextItem()
}

func (p *parser) run() {
	for state := parseStart; state != nil; {
		state = state(p)
	}
	close(p.tokens)
}

// parseStart scans for either an identifier, or an array index
func parseStart(p *parser) parseStateFn {
	for p.peek().typ != itemEOF {
		switch p.peek().typ {
		case itemLeftDelim:
			return parseLeftDelim
		case itemText:
			return parseText
		default:
			return p.errorf("expected text or left delim")
		}
	}
	p.emit(token{tokenEnd, ""})
	return nil
}

// parseText scans for left delim
func parseText(p *parser) parseStateFn {
	i := p.getBufOrNext()
	if i.typ == itemText {
		p.emit(token{tokenText, i.val})
		switch p.peek().typ {
		case itemEOF:
			return parseStart
		default:
			return parseLeftDelim
		}
	}
	return p.errorf("expected text")
}

// parseLeftDelim scans for left delim
func parseLeftDelim(p *parser) parseStateFn {
	if p.getBufOrNext().typ == itemLeftDelim {
		return parseIdentifier
	}
	return p.errorf("expected identifier")
}

// parseIdentifier scans for identifier
func parseIdentifier(p *parser) parseStateFn {
	i := p.getBufOrNext()
	if i.typ == itemIdentifier {
		p.emit(token{tokenIdentifier, i.val})
		return parseRightDelim
	}
	return p.errorf("expected identifier")
}

// parseRightDelim scans for left delim
func parseRightDelim(p *parser) parseStateFn {
	if p.getBufOrNext().typ == itemRightDelim {
		return parseStart
	}
	return p.errorf("expected right delim")
}
