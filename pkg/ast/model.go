package ast

import (
	"fmt"

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
			return fmt.Errorf("unexected token %s, expecting either model or view stanza or '}'", lit)
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
	elements []Element
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

		case CLOSING_BRACE:
			closed = true

		default:
			return fmt.Errorf("unexected token %s, expecting '}'", lit)
		}

		if err != nil {
			return err
		}

	}

	return nil
}

type ElementType int32

const (
	SoftwareSystem ElementType = iota
)

type Element struct {
	ElementType ElementType
	Name        string
}
