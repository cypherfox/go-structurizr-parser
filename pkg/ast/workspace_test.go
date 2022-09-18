package ast_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
	"github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type stmtFactory func() *ast.WorkspaceStatement

type testCase struct {
	label    string
	s        string
	stmt     *ast.WorkspaceStatement // literal datastructure to check
	stmt_fnc stmtFactory             // function to generate data structure to check
	err      string
}

// Ensure the scanner can scan tokens correctly.
func TestWorkspaceParse(t *testing.T) {

	var tests = []testCase{
		{
			label: "Single field statement",
			s: `workspace {
				}`,
			stmt: ast.NewWorkspaceStatement(),
		},

		{
			label: "add empty model",
			s: `workspace {
				model {
				}
			}`,
			stmt_fnc: minimalModelGen,
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
			label: "single software system, no description",
			s: `workspace {
						model {
							softwaresystem app {

							}
						}
					}`,
			stmt_fnc: singleSoftwareSystemStmtNoDescGen,
		},

		{
			label: "single software system with description",
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
			err:   `<string input>:0: found '<end of file>', expected 'workspace'`,
		},
		{
			label: "missing closing brace for model definition",
			s: `workspace {
					model {
						softwaresystem foo {
						}
				`,
			err: `<string input>:8: unexpected end of file in definition of model`,
		},
	}

	runTests(t, tests)
}

func runTests(t *testing.T, tests []testCase) {
	for i, tt := range tests {
		parser := parser.NewParser(strings.NewReader(tt.s), "<string input>")

		stmt, err := ast.Parse(parser)

		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. (%s) %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.label, tt.s, tt.err, err)

		} else if tt.err == "" && tt.stmt != nil && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. (%s) %q\n\nstmt mismatch:\n\nexp=%s\n\ngot=%s\n\n", i, tt.label, tt.s, prettyPrint(tt.stmt), prettyPrint(stmt))

		} else if tt.stmt_fnc != nil && !reflect.DeepEqual(tt.stmt_fnc(), stmt) {
			t.Errorf("%d. (%s) %q\n\nstmt mismatch:\n\nexp=%s\n\ngot=%s\n\n", i, tt.label, tt.s, prettyPrint(tt.stmt_fnc()), prettyPrint(stmt))
		}
	}

}

func prettyPrint(v interface{}) string {
	ret, _ := json.MarshalIndent(v, "", " ")
	return string(ret)
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func minimalModelGen() *ast.WorkspaceStatement {
	return &ast.WorkspaceStatement{
		Model: ast.NewModelStatement(),
	}
}

func singleSoftwareSystemStmtGen() *ast.WorkspaceStatement {
	model := ast.NewModelStatement()

	add := ast.NewSoftwareSystemStatement()

	add.Name = "app"
	add.Description = "This is a multi word description"
	model.AddElement(add)

	ret := ast.NewWorkspaceStatement()
	ret.Model = model

	return ret
}

func singleSoftwareSystemStmtNoDescGen() *ast.WorkspaceStatement {
	ret := singleSoftwareSystemStmtGen()
	swss := ret.Model.GetElementByName("app").(*ast.SoftwareSystemStatement)
	swss.Description = ""

	return ret
}

func singleSoftwareSystemStmtWithTagsGen() *ast.WorkspaceStatement {
	ret := singleSoftwareSystemStmtGen()
	ret.Model.GetElementByName("app").AddTags("tag1", "tag2", "tag3")
	return ret
}
