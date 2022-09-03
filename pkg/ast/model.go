package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type WorkspaceStatement struct {
	Model *ModelStatement
	Views *ViewsStatement
}

func Parse(p *Parser) (*WorkspaceStatement, error) {
	stmt := &WorkspaceStatement{}

	err := stmt.Parse(p)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (w *WorkspaceStatement) Parse(p *Parser) error {

	_, err := p.Expect(WORKSPACE)
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
		case MODEL:
			w.Model = &ModelStatement{}
			err = nextParse(w.Model, p)

		case VIEWS:
			w.Views = &ViewsStatement{}
			err = nextParse(w.Views, p)

		case CLOSING_BRACE:
			closed = true

		default:
			return FmtErrorf(p, "unexpected token %s, expecting either model or views stanza or '}'", lit)
		}

		if err != nil {
			return err
		}

	}

	_, err = p.Expect(EOF)
	if err != nil {
		return err
	}

	return nil
}

func nextParse(stmnt Statement, p *Parser) error {
	return stmnt.Parse(p)
}

type ModelStatement struct {
	// do not use a map, as the name of an object may change, and would not be updated here.
	Elements []Element
}

func (m *ModelStatement) Parse(p *Parser) error {

	// Model was already eaten by the workspace

	_, err := p.Expect(OPEN_BRACE)
	if err != nil {
		return err
	}

	closed := false

	for !closed {
		tok, lit := p.ScanIgnoreWhitespace()
		switch tok {

		case SOFTWARE_SYSTEM:
			e := &SoftwareSystemStatement{}
			m.AddElement(e)
			err = nextParse(e, p)
			if err != nil {
				return err
			}

		case CLOSING_BRACE:
			closed = true

		default:
			return FmtErrorf(p, "found %s ('%s'), expected '}'", tok.String(), lit)
		}

		if err != nil {
			return err
		}

	}

	return nil
}

func (m *ModelStatement) AddElement(e Element) {
	m.Elements = append(m.Elements, e)
}
func (m *ModelStatement) GetElementByName(name string) Element {
	for _, e := range m.Elements {
		if e.GetName() == name {
			return e
		}
	}
	return nil
}
