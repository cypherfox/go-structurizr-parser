package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type DeploymentGroupStatement struct {
	BaseStatement
	Name     string
	Elements []ElementI
}

func NewDeploymentGroupStatement() *DeploymentGroupStatement {
	ret := &DeploymentGroupStatement{}

	ret.AddTags("Element", "DeploymentGroup")

	return ret
}

func (ps *DeploymentGroupStatement) Parse(p *Parser) error {
	lit, err := p.Expect(DEPLOYMENT_GROUP)
	if err != nil {
		return err
	}

	lit, err = p.Expect(IDENTIFIER)
	if err != nil {
		return err
	}
	ps.Name = lit

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

func (p *DeploymentGroupStatement) GetElementType() ElementType {
	return DeploymentGroup
}

func (p *DeploymentGroupStatement) GetName() string {
	return p.Name
}

func (p *DeploymentGroupStatement) AddTags(tags ...string) error {
	return fmt.Errorf("cannot set tags on DeploymentGroup statement, as there are neither tags in the header nor are they allowed as children of the element.")
}

func (dg *DeploymentGroupStatement) AddElement(e ElementI) error {
	eType := e.GetElementType()
	switch eType {
	case DeploymentGroup, DeploymentNode:
		break
	default:
		return fmt.Errorf("element type %s not allowed in deploymentGroup statement", eType.String())
	}

	dg.Elements = append(dg.Elements, e)
	return nil
}

func (dg *DeploymentGroupStatement) GetElementByName(name string) ElementI {
	return GetElementByName(name, dg.Elements)
}
