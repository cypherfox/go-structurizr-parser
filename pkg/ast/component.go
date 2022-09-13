package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type ComponentStatement struct {
	Name        string
	Description string
	Technology  string
	Tags        []string
	Properties  map[string]string
	Elements    []Element
}

func NewComponentStatement() *ComponentStatement {
	ret := &ComponentStatement{}

	ret.AddTags("Element", "Component")

	return ret
}

func (c *ComponentStatement) Parse(p *Parser) error {
	lit, err := p.Expect(COMPONENT)
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

func (p *ComponentStatement) GetElementType() ElementType {
	return Component
}

func (p *ComponentStatement) GetName() string {
	return p.Name
}

func (p *ComponentStatement) AddTags(tags ...string) error {
	p.Tags = append(p.Tags, tags...)
	return nil
}
