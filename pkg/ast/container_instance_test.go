package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestContainerInstanceParse(t *testing.T) {

	/* TODO: activate me.
	currently this does not work

	var tests = []testCase{
		{
			label: "container instance with valid identifier",
			s: `workspace {
					model {
						app = softwareSystem "App" {
							toolshed = container "Toolshed" {}
						}
						deploymentEnvironment {
							deploymentNode "K8s-1" {
								containerInstance toolshed {}
							}
						}
					}
				}`,
			stmt_fnc: containerInstanceGen,
		},

		// ERRORS
		{
			label: "container instance with valid identifier",
			s: `workspace {
					model {
						app = softwareSystem "App" {
							toolshed = container "Toolshed" {}
						}
						deploymentEnvironment {
							deploymentNode "K8s-1" {
								containerInstance numpty {}
							}
						}
					}
				}`,
			err: `<string input>:8: 'numpty' is not a valid identifier for a container`,
		},
		// (none at the moment)
	}

	runTests(t, tests)

	*/
}

func containerInstanceGen() *ast.WorkspaceStatement {
	ret := minimalContainerGen()

	container, _ := ast.WalkPath(ret, "App", "Toolshed")

	component := ast.NewComponentStatement()
	component.Name = "builder1"

	container.(*ast.ContainerStatement).AddElement(component)

	return ret
}
