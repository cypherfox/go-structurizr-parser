package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestGroupParse(t *testing.T) {

	var tests = []testCase{
		{
			label: "minimal group",
			s: `workspace {
					model {
						group "Grp1" {}
					}
				}`,
			stmt_fnc: minimalGroupGen,
		},

		{
			label: "group in enterprise",
			s: `workspace {
					model {
						enterprise "Corp" {
							group "Grp1" {}
						}
					}
				}`,
			stmt_fnc: enterpriseWithGroupGen,
		},

		// ERRORS
		{
			label: "no nested groups",
			s: `workspace {
					model {
						group "Grp1" {
							group "Grp2" {}
						}
					}
				}`,
			err: "<string input>:6: Groups may not be nested",
		},
	}

	runTests(t, tests)
}

func minimalGroupGen() *ast.WorkspaceStatement {
	group := ast.NewGroupStatement()

	group.Parent = ast.Model
	group.Name = "Grp1"

	model := ast.NewModelStatement()
	model.AddElement(group)

	ret := ast.NewWorkspaceStatement()
	ret.Model = model

	return ret
}

func enterpriseWithGroupGen() *ast.WorkspaceStatement {
	ret := minimalEnterpriseGen()

	group := ast.NewGroupStatement()
	group.Parent = ast.Enterprise
	group.Name = "Grp1"

	ret.Model.Enterprise.AddElement(group)

	return ret
}
