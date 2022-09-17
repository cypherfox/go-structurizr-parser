package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type ModelStatement struct {
	// do not use a map, as the name of an object may change, and would not be updated here.
	Elements   []Element
	Enterprise *EnterpriseStatement // there at most be one enterprise be defined per modell
}

func NewModelStatement() *ModelStatement {
	ret := &ModelStatement{}

	return ret
}

func (m *ModelStatement) Parse(p *Parser) error {

	_, err := p.Expect(MODEL)
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

		case ENTERPRISE:
			p.UnScan()
			e := NewEnterpriseStatement()
			if m.Enterprise == nil {
				m.Enterprise = e
			} else {
				return FmtErrorf(p, "only one enterprise per model allowed")
			}
			m.AddElement(e)
			err = nextParse(e, p)

		case PERSON:
			p.UnScan()
			ps := NewPersonStatement()
			m.AddElement(ps)
			err = nextParse(ps, p)

		case GROUP:
			p.UnScan()
			g := NewGroupStatement(Model)
			m.AddElement(g)
			err = nextParse(g, p)

		case SOFTWARE_SYSTEM:
			p.UnScan()
			s := NewSoftwareSystemStatement()
			m.AddElement(s)
			err = nextParse(s, p)

		case DEPLOYMENT_ENV:
			p.UnScan()
			de := NewDeploymentEnvironmentStatement()
			m.AddElement(de)
			err = nextParse(de, p)

		case CONTAINER, COMPONENT, DEPLOYMENT_GROUP, DEPLOYMENT_NODE:
			err = FmtErrorf(p, "found %s, not a valid child for model", tok.String())

		case CLOSING_BRACE:
			closed = true
			continue

		default:
			err = FmtErrorf(p, "found %s ('%s'), expected '}'", tok.String(), lit)
		}

		if err != nil {
			return err
		}

	}

	return nil
}

func (m *ModelStatement) AddElement(e Element) error {
	m.Elements = append(m.Elements, e)
	return nil
}

func (m *ModelStatement) GetElementByName(name string) Element {
	return GetElementByName(name, m.Elements)
}
