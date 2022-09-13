package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestDeploymentEnvironmentParse(t *testing.T) {

	var tests = []testCase{
		{
			label: "minimal deployment environment example",
			s: `workspace {
					model {
						deploymentEnvironment "Testing" {}
					}
				}`,
			stmt_fnc: minimalDeploymentEnvGen,
		},

		// ERRORS
	}

	runTests(t, tests)
}

func minimalDeploymentEnvGen() *ast.WorkspaceStatement {
	ws := minimalModelGen()

	deplEnv := ast.NewDeploymentEnvironmentStatement()
	deplEnv.Name = "Testing"

	ws.Model.AddElement(deplEnv)

	return ws
}
