package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type DeploymentNodeStatement struct {
	BaseStatement
	Name        string
	Description string
	Tags        []string
	Properties  map[string]string
	Elements    []ElementI
}

func NewDeploymentNodeStatement() *DeploymentNodeStatement {
	ret := &DeploymentNodeStatement{}

	ret.AddTags("Element", "DeploymentNode")

	return ret
}

var (
	deploymentNodeAllowedChildren []ElementType = []ElementType{DeploymentNode, InfrastructureNode, SoftwareSystemInstance, ContainerInstance}
)

func (ps *DeploymentNodeStatement) Parse(p *Parser) error {
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

func (p *DeploymentNodeStatement) GetElementType() ElementType {
	return DeploymentNode
}

func (p *DeploymentNodeStatement) GetName() string {
	return p.Name
}

func (p *DeploymentNodeStatement) AddTags(tags ...string) error {
	p.Tags = append(p.Tags, tags...)
	return nil
}

func (dn *DeploymentNodeStatement) AddElement(e ElementI) error {
	var err error

	dn.Elements, err = addAllowedElement(dn.Elements, dn.GetElementType(), e, deploymentNodeAllowedChildren...)

	return err
}

func (dn *DeploymentNodeStatement) GetElementByName(name string) ElementI {
	return GetElementByName(name, dn.Elements)
}
