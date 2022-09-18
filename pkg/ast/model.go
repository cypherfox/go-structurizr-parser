package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type ModelStatement struct {
	BaseStatement
	BaseElementContainer
	Enterprise *EnterpriseStatement // there at most be one enterprise be defined per model
}

func NewModelStatement() *ModelStatement {
	ret := &ModelStatement{
		BaseElementContainer: NewBaseElementContainer(),
	}

	return ret
}

var (
	modelAllowedChildren []ElementType = []ElementType{Enterprise, Group, Person, SoftwareSystem, DeploymentEnvironment, Element}
	modelParseTable      map[Token]ParseFunction
)

func (m *ModelStatement) Parse(p *Parser) error {
	if modelParseTable == nil {
		modelParseTable = setupModelParseTable()
	}

	_, err := p.Expect(MODEL)
	if err != nil {
		return err
	}

	_, err = p.Expect(OPEN_BRACE)
	if err != nil {
		return err
	}

	m.SetBodyOpen(true)

	for m.BodyIsOpen() {
		tok, lit := p.ScanIgnoreWhitespace()
		err = modelParseTable[tok](p, m, tok, lit)
		if err != nil {
			return err
		}

	}

	return nil
}

func (m *ModelStatement) GetElementType() ElementType {
	return Model
}

func (m *ModelStatement) GetName() string {
	// the model has no name
	return ""
}

func (m *ModelStatement) AddTags(tags ...string) error {
	return fmt.Errorf("cannot set tags on Model statement, as there are neither tags in the header nor are they allowed as children of the element.")
}

func setupModelParseTable() map[Token]ParseFunction {
	table := GetDefaultParseFunctionMap()

	table[ENTERPRISE] = enterpriseParse
	table[PERSON] = elementParse
	table[GROUP] = elementParse
	table[SOFTWARE_SYSTEM] = elementParse
	table[DEPLOYMENT_ENV] = elementParse
	table[ELEMENT] = elementParse
	table[IDENTIFIER] = identifierParse

	return table
}

func enterpriseParse(p *Parser, parent interface{}, token Token, literal string) error {
	p.UnScan()
	model := parent.(*ModelStatement)

	if model.Enterprise != nil {
		return FmtErrorf(p, "only one enterprise per model allowed")
	}

	e := NewEnterpriseStatement()
	model.Enterprise = e
	model.AddElement(e)

	return nextParse(e, p)
}

func identifierParse(p *Parser, parent interface{}, token Token, literal string) error {
	_, err := p.Expect(EQUAL)
	if err != nil {
		return err
	}
	parent.(ElementContainer).SetPendingIdentifier(literal)

	return nil
}
