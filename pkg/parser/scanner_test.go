package parser_test

import (
	"strings"
	"testing"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: EOF},
		{s: `#`, tok: ILLEGAL, lit: `#`},
		{s: ` `, tok: WS, lit: " "},
		{s: "\t", tok: WS, lit: "\t"},
		{s: "\n", tok: WS, lit: "\n"},

		// Misc characters
		{s: `,`, tok: COMMA, lit: ","},

		// Identifiers
		{s: `foo`, tok: IDENTIFIER, lit: `foo`},
		{s: `Zx12_3U_-`, tok: IDENTIFIER, lit: `Zx12_3U_`},
		{s: `"a quoted identifier"`, tok: IDENTIFIER, lit: `a quoted identifier`},

		// Keywords
		{s: `workspace`, tok: WORKSPACE, lit: "workspace"},
		{s: `workSpace`, tok: WORKSPACE, lit: "workSpace"},
		{s: `"workSpace"`, tok: WORKSPACE, lit: "workSpace"},
		{s: `model`, tok: MODEL, lit: "model"},
		{s: `views`, tok: VIEWS, lit: "views"},
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.s), "<string input>")
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
