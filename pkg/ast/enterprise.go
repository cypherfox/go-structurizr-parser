package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type EnterpriseStatement struct {
	Name     string
	Elements []Element
}

func NewEnterpriseStatement() *EnterpriseStatement {
	ret := &EnterpriseStatement{}

	return ret
}

func (e *EnterpriseStatement) GetElementType() ElementType {
	return Enterprise
}

func (e *EnterpriseStatement) GetName() string {
	return e.Name
}

func (e *EnterpriseStatement) AddTags(tags ...string) error {
	return fmt.Errorf("cannot set tags on Enterprise statement, as there are neither tags in the header nor are they allowed as children of the element.")
}

func (es *EnterpriseStatement) AddElement(e Element) error {
	etype := e.GetElementType()
	switch etype {
	case Group:
		break
	case Person:
		break
	case SoftwareSystem:
		break
	default:
		return fmt.Errorf("element type %s not allowed in enterprise statement", etype.String())
	}

	es.Elements = append(es.Elements, e)
	return nil
}

func (e *EnterpriseStatement) Parse(p *Parser) error {
	lit, err := p.Expect(ENTERPRISE)
	if err != nil {
		return err
	}

	lit, err = p.Expect(IDENTIFIER)
	if err != nil {
		return err
	}
	e.Name = lit
	_, err = p.Expect(OPEN_BRACE)
	if err != nil {
		return err
	}

	closed := false

	for !closed {
		tok, lit := p.ScanIgnoreWhitespace()
		switch tok {

		case GROUP:
			p.UnScan()
			g := NewGroupStatement()
			g.Parent = Enterprise
			e.AddElement(g)
			err = nextParse(g, p)

			// TODO: PERSON SOFTWARE_SYSTEM RELATIONSHIP
		case PERSON:
			p.UnScan()
			ps := NewPersonStatement()
			e.AddElement(ps)
			err = nextParse(ps, p)

		case CLOSING_BRACE:
			closed = true
			continue

		default:
			err = FmtErrorf(p, "unexpected token %s, expecting '}'", lit)
		}

		if err != nil {
			return err
		}

	}

	return nil

}
