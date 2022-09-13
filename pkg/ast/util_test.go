package ast_test

import (
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
)

func TestWalkPath(t *testing.T) {
	ret := containerWithGroupGen()

	group, err := ast.WalkPath(ret, "Toolshed", "Grp1")
	if err != nil {
		t.Errorf("failed to walk path: %s", err)
		return
	}
	if group == nil {
		t.Errorf("failed to walk path: %s", err)
		return
	}
}
