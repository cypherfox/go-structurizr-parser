package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type DeploymentEnvironmentStatement struct {
	BaseStatement
	Name string
}

func NewDeploymentEnvironmentStatement() *DeploymentEnvironmentStatement {
	ret := &DeploymentEnvironmentStatement{}

	return ret
}

func (des *DeploymentEnvironmentStatement) Parse(p *Parser) error {
	lit, err := p.Expect(DEPLOYMENT_ENV)
	if err != nil {
		return err
	}

	lit, err = p.Expect(IDENTIFIER)
	if err != nil {
		return err
	}
	des.Name = lit

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

func (des *DeploymentEnvironmentStatement) GetElementType() ElementType {
	return DeploymentEnvironment
}

func (des *DeploymentEnvironmentStatement) GetName() string {
	return des.Name
}

func (des *DeploymentEnvironmentStatement) AddTags(tags ...string) error {
	return fmt.Errorf("cannot set tags on DeploymentEnvironment statement, as there are neither tags in the header nor are they allowed as children of the element.")
}
