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
						softwareSystem "App" {
							container "Toolshed" {}
						}
					}
				}`,
			stmt_fnc: minimalContainerGen,
		},

		{
			label: "container with technology example",
			s: `workspace {
					model {
						softwareSystem "App" {
							container "Toolshed" "a set of commonly used tools" "docker" {}
						}
					}
				}`,
			stmt_fnc: containerWithTechnologyGen,
		},

		{
			label: "container with group example",
			s: `workspace {
					model {
						softwareSystem "App" {
							container "Toolshed" {
								group "Grp1" {}
							}
						}
					}
				}`,
			stmt_fnc: containerWithGroupGen,
		},

		{
			label: "container with component example",
			s: `workspace {
					model {
						softwareSystem "App" {
							container "Toolshed" {
								component "builder1" {}
							}
						}
					}
				}`,
			stmt_fnc: containerWithComponentGen,
		},

		{
			label: "container in enterprise",
			s: `workspace {
					model {
						enterprise "Corp" {
							softwareSystem "App" {
								container "Alpine" {}
							}
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
	ws := minimalModelGen()

	soft := ast.NewSoftwareSystemStatement()
	soft.Name = "App"

	container := ast.NewContainerStatement()
	container.Name = "Toolshed"

	soft.AddElement(container)
	ws.Model.AddElement(soft)

	return ws
}

func containerWithTechnologyGen() *ast.WorkspaceStatement {
	ret := minimalContainerGen()

	container, err := ast.WalkPath(ret, "App", "Toolshed")
	if err != nil {
		return nil
	}
	container.(*ast.ContainerStatement).Description = "a set of commonly used tools"
	container.(*ast.ContainerStatement).Technology = "docker"

	return ret
}

func containerWithGroupGen() *ast.WorkspaceStatement {
	ret := minimalContainerGen()

	container, err := ast.WalkPath(ret, "App", "Toolshed")
	if err != nil {
		return nil
	}

	group := ast.NewGroupStatement(ast.Container)
	group.Name = "Grp1"

	container.(*ast.ContainerStatement).AddElement(group)

	return ret
}

func enterpriseWithContainerGen() *ast.WorkspaceStatement {
	ret := enterpriseWithSoftwareSystemGen()

	sSys, _ := ast.WalkPath(ret, "Corp", "App")
	container := ast.NewContainerStatement()
	container.Name = "Alpine"

	sSys.(*ast.SoftwareSystemStatement).AddElement(container)

	return ret
}
