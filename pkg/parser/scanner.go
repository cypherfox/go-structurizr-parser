package parser

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

func isQuote(ch rune) bool {
	return ch == '"'
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isNewLine(ch rune) bool {
	return ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

var eof = rune(0)

// Scanner represents a lexical scanner.
type Scanner struct {
	r      *bufio.Reader
	source string
	line   uint32
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader, s string) *Scanner {
	return &Scanner{
		r:      bufio.NewReader(r),
		source: s,
	}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if isNewLine(ch) {
		s.line++
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if isQuote(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case ',':
		return COMMA, string(ch)
	case '{':
		return OPEN_BRACE, string(ch)
	case '}':
		return CLOSING_BRACE, string(ch)
	case '=':
		return EQUAL, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	ch := s.read()
	quoted := false

	if isQuote(ch) {
		quoted = true
	} else {
		buf.WriteRune(ch)
	}

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		ch = s.read()
		if ch == eof {
			break
		} else if quoted {
			if isQuote(ch) {
				break
			} else { // only another quote will stop it.
				_, _ = buf.WriteRune(ch)
			}
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	token, ok := keywords[strings.ToLower(buf.String())]
	if ok {
		return token, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENTIFIER, buf.String()
}
