package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestContainerParse(t *testing.T) {

	var tests = []testCase{
		{
			label: "minimal container example",
			s: `workspace {
					model {
						container "Toolshed" {}
					}
				}`,
			stmt_fnc: minimalContainerGen,
		},

		{
			label: "container with technology example",
			s: `workspace {
					model {
						container "Toolshed" "a set of commonly used tools" "docker" {}
					}
				}`,
			stmt_fnc: containerWithTechnologyGen,
		},

		{
			label: "container in enterprise",
			s: `workspace {
					model {
						enterprise "Corp" {
							container "Alpine" {}
						}
					}
				}`,
			stmt_fnc: enterpriseWithContainerGen,
		},

		// ERRORS
	}

	runTests(t, tests)
}

func minimalContainerGen() *ast.WorkspaceStatement {
	container := ast.NewContainerStatement()
	container.Name = "Toolshed"

	model := ast.NewModelStatement()
	model.AddElement(container)

	ret := ast.NewWorkspaceStatement()
	ret.Model = model

	return ret
}

func containerWithTechnologyGen() *ast.WorkspaceStatement {
	ret := minimalContainerGen()

	container := ret.Model.GetElementByName("Toolshed")
	container.(*ast.ContainerStatement).Description = "a set of commonly used tools"
	container.(*ast.ContainerStatement).Technology = "docker"

	return ret
}

func enterpriseWithContainerGen() *ast.WorkspaceStatement {
	ret := minimalEnterpriseGen()

	container := ast.NewContainerStatement()
	container.Name = "Toolshed"

	ret.Model.Enterprise.AddElement(container)

	return ret
}
