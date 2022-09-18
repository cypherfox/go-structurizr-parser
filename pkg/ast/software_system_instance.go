package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type SoftwareSystemInstanceStatement struct {
	BaseStatement
	Name        string
	Description string
	Tags        []string
	Properties  map[string]string
	Elements    []ElementI
}

func NewSoftwareSystemInstanceStatement() *SoftwareSystemInstanceStatement {
	ret := &SoftwareSystemInstanceStatement{}

	ret.AddTags("Element", "SoftwareSystemInstance")

	return ret
}

func (ps *SoftwareSystemInstanceStatement) Parse(p *Parser) error {
	lit, err := p.Expect(DEPLOYMENT_NODE)
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

func (p *SoftwareSystemInstanceStatement) GetElementType() ElementType {
	return SoftwareSystemInstance
}

func (p *SoftwareSystemInstanceStatement) GetName() string {
	return p.Name
}

func (p *SoftwareSystemInstanceStatement) AddTags(tags ...string) error {
	p.Tags = append(p.Tags, tags...)
	return nil
}

func (dn *SoftwareSystemInstanceStatement) AddElement(e ElementI) error {
	eType := e.GetElementType()
	switch eType {
	case SoftwareSystemInstance, InfrastructureNode, ContainerInstance:
		break
	default:
		return fmt.Errorf("element type %s not allowed in SoftwareSystemInstance statement", eType.String())
	}

	dn.Elements = append(dn.Elements, e)
	return nil
}

func (dn *SoftwareSystemInstanceStatement) GetElementByName(name string) ElementI {
	return GetElementByName(name, dn.Elements)
}
