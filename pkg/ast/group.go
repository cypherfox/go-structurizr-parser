package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type GroupStatement struct {
	Parent   ElementType
	Name     string
	Elements []Element
}

func NewGroupStatement(parent ElementType) *GroupStatement {
	ret := &GroupStatement{
		Parent: parent,
	}

	return ret
}

func (g *GroupStatement) GetElementType() ElementType {
	return Group
}

func (g *GroupStatement) GetName() string {
	return g.Name
}

func (g *GroupStatement) AddTags(tags ...string) error {
	return fmt.Errorf("cannot set tags on Group statement, as there are neither tags in the header nor are they allowed as children of the element.")

}

func (g *GroupStatement) Parse(p *Parser) error {
	lit, err := p.Expect(GROUP)
	if err != nil {
		return err
	}

	lit, err = p.Expect(IDENTIFIER)
	if err != nil {
		return err
	}
	g.Name = lit

	_, err = p.Expect(OPEN_BRACE)
	if err != nil {
		return err
	}

	closed := false

	for !closed {
		tok, lit := p.ScanIgnoreWhitespace()
		switch tok {

		case COMPONENT:
			if g.Parent == Container {
				p.UnScan()
				c := NewComponentStatement()
				g.AddElement(c)
				err = nextParse(c, p)
			} else {
				err = FmtErrorf(p, "component is onyl allowed in group as child of container")
			}

		case CLOSING_BRACE:
			closed = true
			continue

		case GROUP:
			err = FmtErrorf(p, "Groups may not be nested")
		default:
			err = FmtErrorf(p, "unexpected token %s, expecting '}'", lit)
		}

		if err != nil {
			return err
		}

	}

	return nil

}

func (g *GroupStatement) GetElementByName(name string) Element {
	return GetElementByName(name, g.Elements)
}

func (g *GroupStatement) AddElement(e Element) {
	g.Elements = append(g.Elements, e)
}
