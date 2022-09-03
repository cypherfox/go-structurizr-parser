package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestEnterpriseParse(t *testing.T) {

	var tests = []testCase{
		{
			label: "minimal enterprise",
			s: `workspace {
					model {
						enterprise "Corp" {}
					}
				}`,
			stmt_fnc: minimalEnterpriseGen,
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

func minimalEnterpriseGen() *ast.WorkspaceStatement {
	enterprise := &ast.EnterpriseStatement{
		Name: "Corp",
	}

	model := &ast.ModelStatement{
		Enterprise: enterprise,
	}

	ret := &ast.WorkspaceStatement{
		Model: model,
	}

	return ret
}
