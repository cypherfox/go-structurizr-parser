package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestComponentParse(t *testing.T) {

	var tests = []testCase{
		{
			label: "container with component example",
			s: `workspace {
					model {
						container "Toolshed" {
							component "builder1" {}
						}
					}
				}`,
			stmt_fnc: containerWithComponentGen,
		},

		{
			label: "container with two component example",
			s: `workspace {
					model {
						container "Toolshed" {
							component "builder1" {}
							component "builder2" {}
						}
					}
				}`,
			stmt_fnc: containerWithTwoComponentGen,
		},

		{
			label: "component in group example",
			s: `workspace {
					model {
						container "Toolshed" {
							group "Grp1" {
								component "builder1" {}
							}
						}
					}
				}`,
			stmt_fnc: componentInGroupGen,
		},

		{
			label: "container with component example",
			s: `workspace {
					model {
						container "Toolshed" {
							component "builder1" {}
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

func containerWithComponentGen() *ast.WorkspaceStatement {
	ret := minimalContainerGen()

	container := ret.Model.GetElementByName("Toolshed")

	component := ast.NewComponentStatement()
	component.Name = "builder1"

	container.(*ast.ContainerStatement).AddElement(component)

	return ret
}

func containerWithTwoComponentGen() *ast.WorkspaceStatement {
	ret := containerWithComponentGen()

	container := ret.Model.GetElementByName("Toolshed")

	component := ast.NewComponentStatement()
	component.Name = "builder2"

	container.(*ast.ContainerStatement).AddElement(component)

	return ret
}

func componentInGroupGen() *ast.WorkspaceStatement {
	ret := containerWithGroupGen()

	group, err := ast.WalkPath(ret, "Toolshed", "Grp1")
	if err != nil {
		return nil
	}

	component := ast.NewComponentStatement()
	component.Name = "builder1"
	group.(ast.ElementContainer).AddElement(component)

	return ret
}
