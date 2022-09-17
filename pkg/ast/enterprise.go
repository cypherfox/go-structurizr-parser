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
			g := NewGroupStatement(Enterprise)
			e.AddElement(g)
			err = nextParse(g, p)

			// TODO: RELATIONSHIP
		case PERSON:
			p.UnScan()
			ps := NewPersonStatement()
			e.AddElement(ps)
			err = nextParse(ps, p)

		case SOFTWARE_SYSTEM:
			p.UnScan()
			s := NewSoftwareSystemStatement()
			e.AddElement(s)
			err = nextParse(s, p)

		case CONTAINER:
			p.UnScan()
			c := NewContainerStatement()
			e.AddElement(c)
			err = nextParse(c, p)

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

func (es *EnterpriseStatement) AddElement(e Element) error {
	eType := e.GetElementType()
	switch eType {
	case Group, Person, SoftwareSystem:
		break
	default:
		return fmt.Errorf("element type %s not allowed in enterprise statement", eType.String())
	}

	es.Elements = append(es.Elements, e)
	return nil
}

func (es *EnterpriseStatement) GetElementByName(name string) Element {
	return GetElementByName(name, es.Elements)
}
