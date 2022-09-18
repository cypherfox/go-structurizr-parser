package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

func GetElementByName(name string, elements []ElementI) ElementI {
	for _, e := range elements {
		if e.GetName() == name {
			return e
		}
	}
	return nil

}

func nextParse(stmnt Statement, p *Parser) error {
	return stmnt.Parse(p)
}

// WalkPath will return an element identified by a list of names.
func WalkPath(ws *WorkspaceStatement, path ...string) (ElementI, error) {
	var elem ElementI

	elem = ws.Model.GetElementByName(path[0])
	if elem == nil {
		return nil, fmt.Errorf("Element by name %s not found.", path[0])
	}
	for _, name := range path[1:] {

		val, ok := elem.(ElementContainer)
		if !ok {
			return nil, fmt.Errorf("type %s does not support querying", elem.GetElementType().String())
		}

		elem = val.GetElementByName(name)
		if elem == nil {
			return nil, fmt.Errorf("Element by name %s not found.", path[0])
		}

	}
	return elem, nil
}

func addAllowedElement(elements []ElementI, parentType ElementType, elem ElementI, allowed ...ElementType) ([]ElementI, error) {
	eType := elem.GetElementType()

	for _, t := range allowed {
		if t == eType {
			elements = append(elements, elem)
			return elements, nil
		}
	}

	return elements, fmt.Errorf("element type %s not allowed in %s statement", eType.String(), parentType.String())
}
