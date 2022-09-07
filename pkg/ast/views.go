package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type ViewsStatement struct {
	name string
}

func NewViewsStatement() *ViewsStatement {
	ret := &ViewsStatement{}

	return ret
}

func (v *ViewsStatement) Parse(p *Parser) error {

	_, err := p.Expect(VIEWS)
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
			return fmt.Errorf("found %s ('%s'), expected '}'", tok.String(), lit)
		}

		if err != nil {
			return err
		}

	}

	return nil
}
