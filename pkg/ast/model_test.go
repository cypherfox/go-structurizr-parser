package ast_test

import (
	"strings"
	"testing"

	"github.com/cypherfox/go-structurizr-parser/pkg/ast"
	"github.com/cypherfox/go-structurizr-parser/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestIdentifierParsing(t *testing.T) {
	assert := assert.New(t)

	input := `workspace {
		model {
			app = softwaresystem "App" {

			}
		}
	}`

	parser := parser.NewParser(strings.NewReader(input), "<string input>")
	stmt, err := ast.Parse(parser)
	assert.Nil(err)

	elem := stmt.Model.GetElementByIdentifier("app")
	assert.NotNil(elem)
	assert.Equal(ast.SoftwareSystem, elem.GetElementType)
}
