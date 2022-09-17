package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type ContainerStatement struct {
	Name        string
	Description string
	Technology  string
	Tags        []string
	Properties  map[string]string
	Elements    []Element
}

func NewContainerStatement() *ContainerStatement {
	ret := &ContainerStatement{}

	ret.AddTags("Element", "Container")

	return ret
}

func (c *ContainerStatement) Parse(p *Parser) error {
	lit, err := p.Expect(CONTAINER)
	if err != nil {
		return err
	}

	lit, err = p.Expect(IDENTIFIER)
	if err != nil {
		return err
	}
	c.Name = lit

	err = p.Maybe(IDENTIFIER, func(tok Token, lit string) error {
		c.Description = lit
		return nil
	})
	if err != nil {
		return err
	}

	err = p.Maybe(IDENTIFIER, func(tok Token, lit string) error {
		c.Technology = lit
		return nil
	})
	if err != nil {
		return err
	}

	pTags, err := p.ParseTags()
	c.AddTags(pTags...)

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
			g := NewGroupStatement(Container)
			c.AddElement(g)
			err = nextParse(g, p)

		case COMPONENT:
			p.UnScan()
			comp := NewComponentStatement()
			c.AddElement(comp)
			err = nextParse(comp, p)

		case CLOSING_BRACE:
			closed = true

		default:
			err = FmtErrorf(p, "unexected token %s, expecting '}'", lit)
		}

		if err != nil {
			return err
		}

	}

	return nil
}

func (c *ContainerStatement) GetElementType() ElementType {
	return Container
}

func (c *ContainerStatement) GetName() string {
	return c.Name
}

func (c *ContainerStatement) AddTags(tags ...string) error {
	c.Tags = append(c.Tags, tags...)
	return nil
}

func (c *ContainerStatement) AddElement(e Element) error {
	c.Elements = append(c.Elements, e)
	return nil
}

func (c *ContainerStatement) GetElementByName(name string) Element {
	return GetElementByName(name, c.Elements)
}
