package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestSoftwareSystemParse(t *testing.T) {

	var tests = []testCase{
		{
			label: "minimal person example",
			s: `workspace {
					model {
						softwareSystem "App" {}
					}
				}`,
			stmt_fnc: minimalSoftwareSystemGen,
		},

		{
			label: "minimal person example",
			s: `workspace {
					model {
						softwareSystem "App" {
							group "Grp1" {}
						}
					}
				}`,
			stmt_fnc: softwareSystemWithGroupGen,
		},

		{
			label: "group in enterprise",
			s: `workspace {
					model {
						enterprise "Corp" {
							softwareSystem "App" {}
						}
					}
				}`,
			stmt_fnc: enterpriseWithSoftwareSystemGen,
		},

		// ERRORS
	}

	runTests(t, tests)
}

func minimalSoftwareSystemGen() *ast.WorkspaceStatement {
	system := ast.NewSoftwareSystemStatement()
	system.Name = "App"

	model := ast.NewModelStatement()
	model.AddElement(system)

	ret := ast.NewWorkspaceStatement()
	ret.Model = model

	return ret
}

func softwareSystemWithGroupGen() *ast.WorkspaceStatement {
	ws := minimalSoftwareSystemGen()

	system := ws.Model.GetElementByName("App")

	group := ast.NewGroupStatement(ast.SoftwareSystem)
	group.Name = "Grp1"

	system.(ast.ElementContainer).AddElement(group)

	return ws
}

func enterpriseWithSoftwareSystemGen() *ast.WorkspaceStatement {
	ret := minimalEnterpriseGen()

	system := ast.NewSoftwareSystemStatement()
	system.Name = "App"

	ret.Model.Enterprise.AddElement(system)

	return ret
}
