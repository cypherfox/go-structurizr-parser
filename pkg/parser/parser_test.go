package parser_test

import (
	"reflect"
	"strings"
	"testing"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

// Ensure the scanner can scan tokens correctly.
func TestParser_ParseTags(t *testing.T) {
	var tests = []struct {
		s    string
		tags []string
		err  string
	}{
		{
			s:    ``,
			tags: []string{},
		},

		{
			s:    `{`,
			tags: []string{},
		},

		{
			s:    `foo {`,
			tags: []string{"foo"},
		},

		{
			s:    `foo bar {`,
			tags: []string{"foo", "bar"},
		},

		{
			s:    `foo,bar {`,
			tags: []string{"foo", "bar"},
		},

		{
			s:    `foo, bar {`,
			tags: []string{"foo", "bar"},
		},

		{
			s:    `"foo" {`,
			tags: []string{"foo"},
		},

		{
			s:    `"foo" "bar" {`,
			tags: []string{"foo", "bar"},
		},

		{
			s:    `"foo","bar" {`,
			tags: []string{"foo", "bar"},
		},

		{
			s:    `"foo", "bar" {`,
			tags: []string{"foo", "bar"},
		},

		{
			s:    `"foo", "bar" "baz" {`,
			tags: []string{"foo", "bar", "baz"},
		},
	}

	for i, tt := range tests {
		p := NewParser(strings.NewReader(tt.s), "<string input>")

		result, err := p.ParseTags()

		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)

		} else if tt.err == "" && tt.tags != nil && !reflect.DeepEqual(tt.tags, result) {
			t.Errorf("%d. %q\n token mismatch: exp=%q got=%q\n\n", i, tt.s, tt.tags, result)

		}

	}

}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
