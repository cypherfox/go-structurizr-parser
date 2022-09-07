package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type PersonStatement struct {
	Name        string
	Description string
	Tags        []string
	Properties  map[string]string
	Elements    []*Element
}

func NewPersonStatement() *PersonStatement {
	ret := &PersonStatement{}

	ret.AddTags("Element", "Person")

	return ret
}

func (ps *PersonStatement) Parse(p *Parser) error {
	lit, err := p.Expect(PERSON)
	if err != nil {
		return err
	}

	lit, err = p.Expect(IDENTIFIER)
	if err != nil {
		return err
	}
	ps.Name = lit

	err = p.Maybe(IDENTIFIER, func(tok Token, lit string) error {
		ps.Description = lit
		return nil
	})
	if err != nil {
		return err
	}

	pTags, err := p.ParseTags()
	ps.AddTags(pTags...)

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

func (p *PersonStatement) GetElementType() ElementType {
	return Person
}

func (p *PersonStatement) GetName() string {
	return p.Name
}

func (p *PersonStatement) AddTags(tags ...string) error {
	p.Tags = append(p.Tags, tags...)
	return nil
}
