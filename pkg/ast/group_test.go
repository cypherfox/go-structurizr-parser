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
	group := &ast.GroupStatement{
		Parent: ast.Model,
		Name:   "Grp1",
	}

	model := &ast.ModelStatement{}
	model.AddElement(group)

	ret := &ast.WorkspaceStatement{
		Model: model,
	}

	return ret
}

func enterpriseWithGroupGen() *ast.WorkspaceStatement {
	ret := minimalEnterpriseGen()

	group := &ast.GroupStatement{
		Parent: ast.Enterprise,
		Name:   "Grp1",
	}

	ret.Model.Enterprise.AddElement(group)

	return ret
}
