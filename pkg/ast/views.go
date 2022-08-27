package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type ViewsStatement struct {
	name string
}

func (v *ViewsStatement) Parse(p *Parser) error {
	v.name = "<not set yet>"
	return nil
}
