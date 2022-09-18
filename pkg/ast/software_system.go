package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type SoftwareSystemStatement struct {
	BaseStatement
	BaseElementContainer
	Name        string
	Description string
	Tags        []string
	Properties  map[string]string
	Elements    []ElementI
}

func NewSoftwareSystemStatement() *SoftwareSystemStatement {
	ret := &SoftwareSystemStatement{}

	ret.AddTags("Element", "Software System")

	return ret
}

func (s *SoftwareSystemStatement) Parse(p *Parser) error {
	lit, err := p.Expect(SOFTWARE_SYSTEM)
	if err != nil {
		return err
	}

	lit, err = p.Expect(IDENTIFIER)
	if err != nil {
		return err
	}
	s.Name = lit

	err = p.Maybe(IDENTIFIER, func(tok Token, lit string) error {
		s.Description = lit
		return nil
	})
	if err != nil {
		return err
	}

	pTags, err := p.ParseTags()
	s.AddTags(pTags...)

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
			g := NewGroupStatement(SoftwareSystem)
			s.AddElement(g)
			err = nextParse(g, p)

		case CONTAINER:
			p.UnScan()
			c := NewContainerStatement()
			s.AddElement(c)
			err = nextParse(c, p)

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

func (s *SoftwareSystemStatement) GetElementType() ElementType {
	return SoftwareSystem
}

func (s *SoftwareSystemStatement) GetName() string { return s.Name }

func (s *SoftwareSystemStatement) AddTags(tags ...string) error {
	s.Tags = append(s.Tags, tags...)
	return nil
}
