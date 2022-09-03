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
			stmt_fnc: minimalEnterpriseGen,
		},

		{
			label: "group in enterprise",
			s: `workspace {
					model {
						enterprise "Corp" {
							group "Grp1 {}
						}
					}
				}`,
			stmt_fnc: enterpriseWithGroupGen,
		},

		// ERRORS
		{
			label: "minimal enterprise",
			s: `workspace {
					model {
						enterprise "Corp1" {}
						enterprise "Corp2" {}
					}
				}`,
			err: "<string input>:6: only one enterprise per model allowed",
		},

		{
			label: "minimal enterprise",
			s: `workspace {
					model {
						enterprise "Corp1" {}
						enterprise "Corp1" {}
					}
				}`,
			err: "<string input>:6: only one enterprise per model allowed",
		},
	}

	runTests(t, tests)
}

func minimalGroupGen() *ast.WorkspaceStatement {
	group := &ast.GroupStatement{
		Name: "Grp1",
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
		Name: "Grp1",
	}

	ret.Model.Enterprise.AddElement(group)
}
