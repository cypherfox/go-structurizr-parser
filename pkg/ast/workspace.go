package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type WorkspaceStatement struct {
	Model *ModelStatement
	Views *ViewsStatement
}

func NewWorkspaceStatement() *WorkspaceStatement {
	ret := &WorkspaceStatement{}

	return ret
}

func Parse(p *Parser) (*WorkspaceStatement, error) {
	stmt := NewWorkspaceStatement()

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
			p.UnScan()
			w.Model = NewModelStatement()
			err = nextParse(w.Model, p)

		case VIEWS:
			p.UnScan()
			w.Views = NewViewsStatement()
			err = nextParse(w.Views, p)

		case CLOSING_BRACE:
			closed = true

		default:
			err = FmtErrorf(p, "unexpected token %s, expecting either model or views stanza or '}'", lit)
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
