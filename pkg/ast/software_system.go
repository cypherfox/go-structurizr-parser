package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type SoftwareSystemStatement struct {
	Name        string
	Description string
	Tags        []string
	Properties  map[string]string
	Elements    []*Element
}

func (s *SoftwareSystemStatement) Parse(p *Parser) error {
	// SoftwareSystem was already eaten by the workspace

	lit, err := p.Expect(IDENTIFIER)
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
			return fmtErrorf(p, "unexected token %s, expecting '}'", lit)
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

func (s *SoftwareSystemStatement) AddTags(tags ...string) { s.Tags = append(s.Tags, tags...) }