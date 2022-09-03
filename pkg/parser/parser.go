package parser

import (
	"fmt"
	"io"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader, input string) *Parser {
	return &Parser{s: NewScanner(r, input)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unScan pushes the previously read token back onto the buffer.
func (p *Parser) unScan() { p.buf.n = 1 }

// ScanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) ScanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	fmt.Printf("tok: %s lit: >%s<\n", tok.String(), lit)
	return
}

// Expect expects a certain token and fail with usefull error message if it is
// not found. It will return the literal if successful
func (p *Parser) Expect(expected Token) (string, error) {
	lit := ""
	var tok Token
	if tok, lit = p.ScanIgnoreWhitespace(); tok != expected {
		return "", FmtErrorf(p, "found '%s', expected '%s'", tok.String(), expected.String())
	}
	return lit, nil
}

// ApplyFkt defines code to be executed if a desired token has been found.
type ApplyFkt func(token Token, literal string) error

// Maybe tries to read a certain token next, but unreads the token when it does not match.
// If the token matches, function apply is executed
func (p *Parser) Maybe(expected Token, apply ApplyFkt) error {
	lit := ""
	var tok Token
	if tok, lit = p.ScanIgnoreWhitespace(); tok == expected {
		apply(expected, lit)
	} else {
		p.unScan()
	}

	return nil
}

// ParseTags will attempt to scan a set of identifiers and quoted strings until either a newline or an opening brace is encountered
func (p *Parser) ParseTags() ([]string, error) {
	retArr := []string{}

	for {
		tok, lit := p.ScanIgnoreWhitespace()

		switch tok {
		case EOF:
			return retArr, nil
		case OPEN_BRACE:
			p.unScan()
			return retArr, nil
		case IDENTIFIER:
			retArr = append(retArr, lit)
		case COMMA:
			// ignore it
		default:
			return nil, fmt.Errorf("unexpected token %s lit:<%s>", tok.String(), lit)
		}
	}
}

func (p *Parser) GetScanLine() uint32 {
	return p.s.line
}

func (p *Parser) GetScanSource() string {
	return p.s.source
}

func FmtErrorf(p *Parser, fmtStr string, args ...interface{}) error {
	posInfo := []interface{}{p.GetScanSource(), p.GetScanLine()}

	allArgs := append(posInfo, args...)
	return fmt.Errorf("%s:%d: "+fmtStr, allArgs...)
}
