package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestPersonParse(t *testing.T) {

	var tests = []testCase{
		{
			label: "minimal person example",
			s: `workspace {
					model {
						person "User" {}
					}
				}`,
			stmt_fnc: minimalPersonGen,
		},

		{
			label: "group in enterprise",
			s: `workspace {
					model {
						enterprise "Corp" {
							person "User1" {}
						}
					}
				}`,
			stmt_fnc: enterpriseWithPersonGen,
		},

		// ERRORS
	}

	runTests(t, tests)
}

func minimalPersonGen() *ast.WorkspaceStatement {
	person := ast.NewPersonStatement()
	person.Name = "User"

	model := ast.NewModelStatement()
	model.AddElement(person)

	ret := ast.NewWorkspaceStatement()
	ret.Model = model

	return ret
}

func enterpriseWithPersonGen() *ast.WorkspaceStatement {
	ret := minimalEnterpriseGen()

	group := ast.NewPersonStatement()
	group.Name = "User1"

	ret.Model.Enterprise.AddElement(group)

	return ret
}
