package ast_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
	"github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type stmtFactory func() *ast.WorkspaceStatement

// Ensure the scanner can scan tokens correctly.
func TestParse(t *testing.T) {

	var tests = []struct {
		label    string
		s        string
		stmt     *ast.WorkspaceStatement // literal datastructure to check
		stmt_fnc stmtFactory             // function to generate data structure to check
		err      string
	}{
		{
			label: "Single field statement",
			s: `workspace {
				}`,
			stmt: &ast.WorkspaceStatement{},
		},

		{
			label: "add empty model",
			s: `workspace {
				model {
				}
			}`,
			stmt: &ast.WorkspaceStatement{
				Model: &ast.ModelStatement{},
			},
		},

		{
			label: "add empty views",
			s: `workspace {
				model {
				}
				views {
				}
			}`,
			stmt: &ast.WorkspaceStatement{
				Model: &ast.ModelStatement{},
				Views: &ast.ViewsStatement{},
			},
		},

		{
			label: "single software system",
			s: `workspace {
						model {
							softwaresystem app "This is a multi word description" {

							}
						}
					}`,
			stmt_fnc: singleSoftwareSystemStmtGen,
		},

		{
			label: "single software system, with tags",
			s: `workspace {
						model {
							softwaresystem app "This is a multi word description" tag1,tag2,tag3 {

							}
						}
					}`,
			stmt_fnc: singleSoftwareSystemStmtWithTagsGen,
		},

		// Errors
		{
			label: "empty file",
			s:     ``,
			err:   `found '<end of file>', expected 'workspace'`,
		},
		{
			label: "missing closing brace for model definition",
			s: `workspace {
					model {
				softwaresystem {

				}
			}
		}`, err: `found '{', expected '<an identifier>'`},
	}

	for i, tt := range tests {
		parser := parser.NewParser(strings.NewReader(tt.s), "<string input>")

		stmt, err := ast.Parse(parser)

		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. (%s) %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.label, tt.s, tt.err, err)

		} else if tt.err == "" && tt.stmt != nil && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. (%s) %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.label, tt.s, tt.stmt, stmt)

		} else if tt.stmt_fnc != nil && !reflect.DeepEqual(tt.stmt_fnc(), stmt) {
			t.Errorf("%d. (%s) %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.label, tt.s, tt.stmt_fnc(), stmt)
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

func singleSoftwareSystemStmtGen() *ast.WorkspaceStatement {
	model := &ast.ModelStatement{
		Elements: []ast.Element{},
	}

	add := &ast.SoftwareSystemStatement{
		Name:        "app",
		Description: "This is a multi word description",
	}
	model.AddElement(add)

	ret := &ast.WorkspaceStatement{
		Model: model,
	}

	return ret
}

func singleSoftwareSystemStmtWithTagsGen() *ast.WorkspaceStatement {
	ret := singleSoftwareSystemStmtGen()
	ret.Model.GetElementByName("app").AddTags("tag1", "tag2", "tag3")
	return ret
}
