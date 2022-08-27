package ast_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
	"github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

// Ensure the scanner can scan tokens correctly.
func TestParse(t *testing.T) {

	var tests = []struct {
		s    string
		stmt *ast.WorkspaceStatement
		err  string
	}{
		// Single field statement
		{
			s: `workspace {
				}`,
			stmt: &ast.WorkspaceStatement{},
		},

		// add empty model
		{
			s: `workspace {
				model {
				}
			}`,
			stmt: &ast.WorkspaceStatement{
				Model: &ast.ModelStatement{},
			},
		},
		// Errors
		{s: ``, err: `found <end of file>, expected workspace`},
	}

	for i, tt := range tests {
		parser := parser.NewParser(strings.NewReader(tt.s))

		stmt, err := ast.Parse(parser)

		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
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
