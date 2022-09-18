package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type InfrastructureNodeStatement struct {
	BaseStatement
	Name        string
	Description string
	Tags        []string
	Properties  map[string]string
}

func NewInfrastructureNodeStatement() *InfrastructureNodeStatement {
	ret := &InfrastructureNodeStatement{}

	ret.AddTags("Element", "InfrastructureNode")

	return ret
}

func (ps *InfrastructureNodeStatement) Parse(p *Parser) error {
	lit, err := p.Expect(INFRASTRUCTURE_NODE)
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
			err = FmtErrorf(p, "unexpected token %s, expecting '}'", lit)
		}

		if err != nil {
			return err
		}

	}

	return nil
}

func (p *InfrastructureNodeStatement) GetElementType() ElementType {
	return InfrastructureNode
}

func (p *InfrastructureNodeStatement) GetName() string {
	return p.Name
}

func (p *InfrastructureNodeStatement) AddTags(tags ...string) error {
	p.Tags = append(p.Tags, tags...)
	return nil
}
