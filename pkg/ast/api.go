package ast

import (
	"fmt"

	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type ElementType Token

const (
	Model                  ElementType = ElementType(MODEL)
	SoftwareSystem         ElementType = ElementType(SOFTWARE_SYSTEM)
	Enterprise             ElementType = ElementType(ENTERPRISE)
	Group                  ElementType = ElementType(GROUP)
	Person                 ElementType = ElementType(PERSON)
	Container              ElementType = ElementType(CONTAINER)
	Component              ElementType = ElementType(COMPONENT)
	DeploymentEnvironment  ElementType = ElementType(DEPLOYMENT_ENV)
	DeploymentGroup        ElementType = ElementType(DEPLOYMENT_GROUP)
	DeploymentNode         ElementType = ElementType(DEPLOYMENT_NODE)
	InfrastructureNode     ElementType = ElementType(INFRASTRUCTURE_NODE)
	SoftwareSystemInstance ElementType = ElementType(SOFTWARE_SYSTEM_INSTANCE)
	ContainerInstance      ElementType = ElementType(CONTAINER_INSTANCE)
	Element                ElementType = ElementType(ELEMENT)
)

func (e ElementType) String() string {
	t := Token(e)
	return t.String()
}

type ElementContainer interface {
	PendingIdentifier() string
	SetPendingIdentifier(string)
	GetElementByName(name string) ElementI
	// will also assign any pending identifiers
	AddElement(ElementI) error
	AddIdentifier(string, ElementI)
	GetElementByIdentifier(string) ElementI
}

type BaseElementContainer struct {
	// do not use a map, as the name of an object may change, and would not be updated here.
	Elements          []ElementI
	identifiers       map[string]ElementI
	pendingIdentifier string
}

func NewBaseElementContainer() BaseElementContainer {
	return BaseElementContainer{
		identifiers: map[string]ElementI{},
	}
}

func (b *BaseElementContainer) PendingIdentifier() string {
	return b.pendingIdentifier
}

func (b *BaseElementContainer) SetPendingIdentifier(id string) {
	b.pendingIdentifier = id
}

func (b *BaseElementContainer) AddElement(e ElementI) error {
	id := b.PendingIdentifier()
	if id != "" {
		b.AddIdentifier(id, e)
	}
	b.Elements = append(b.Elements, e)
	return nil

}

func (b *BaseElementContainer) AddIdentifier(id string, elem ElementI) {
	// TODO: fix thread safety
	b.identifiers[id] = elem
	b.pendingIdentifier = ""
}

func (b *BaseElementContainer) GetElementByName(name string) ElementI {
	return GetElementByName(name, b.Elements)
}

func (b *BaseElementContainer) GetElementByIdentifier(id string) ElementI {
	return b.identifiers[id]
}

type ElementI interface {
	GetElementType() ElementType
	GetName() string
	AddTags(tags ...string) error
}

type Statement interface {
	Parse(p *Parser) error
	SetBodyOpen(open bool)
	BodyIsOpen() bool
}

type BaseStatement struct {
	isOpen bool
}

func (b *BaseStatement) SetBodyOpen(open bool) {
	b.isOpen = open
}

func (b *BaseStatement) BodyIsOpen() bool {
	return b.isOpen
}

type ParseFunction func(*Parser, interface{}, Token, string) error

func GetDefaultParseFunctionMap() map[Token]ParseFunction {
	return map[Token]ParseFunction{
		ILLEGAL: notImplementedParse,
		EOF:     unexpectedEOFParse,
		WS:      notImplementedParse,

		// Literals
		IDENTIFIER: notImplementedParse,

		// MISC chars
		COMMA:         notImplementedParse,
		OPEN_BRACE:    notImplementedParse,
		CLOSING_BRACE: closeBodyParse,
		EQUAL:         notImplementedParse,
		QUOTE:         notImplementedParse,

		RELATION: notImplementedParse,

		// Model Keywords
		WORKSPACE:                notValidChildParse,
		MODEL:                    notValidChildParse,
		ENTERPRISE:               notValidChildParse,
		GROUP:                    notValidChildParse,
		PERSON:                   notValidChildParse,
		SOFTWARE_SYSTEM:          notValidChildParse,
		CONTAINER:                notValidChildParse,
		COMPONENT:                notValidChildParse,
		DEPLOYMENT_ENV:           notValidChildParse,
		DEPLOYMENT_GROUP:         notValidChildParse,
		DEPLOYMENT_NODE:          notValidChildParse,
		INFRASTRUCTURE_NODE:      notValidChildParse,
		SOFTWARE_SYSTEM_INSTANCE: notValidChildParse,
		CONTAINER_INSTANCE:       notValidChildParse,
		ELEMENT:                  notValidChildParse,

		// Views Keywords
		VIEWS:            notValidChildParse,
		SYSTEM_LANDSCAPE: notValidChildParse,
		SYSTEM_CONTEXT:   notValidChildParse,
		FILTERED:         notValidChildParse,
		DYNAMIC:          notValidChildParse,
		DEPLOYMENT:       notValidChildParse,
		CUSTOM:           notValidChildParse,
		STYLES:           notValidChildParse,
		THEME:            notValidChildParse,
		THEMES:           notValidChildParse,
		BRANDING:         notValidChildParse,
		TERMINOLOGY:      notValidChildParse,

		CONFIGURATION: notImplementedParse,
		USERS:         notImplementedParse,

		// Reference keywords and pragmas
		BANG_DOCS:                  notValidChildParse,
		BANG_ADRS:                  notValidChildParse,
		BANG_IDENTIFIERS:           notValidChildParse,
		BANG_IMPLIED_RELATIONSHIPS: notValidChildParse,
		BANG_REF:                   notValidChildParse,
		BANG_INCLUDE:               notValidChildParse,
		EXTENDS:                    notValidChildParse,
	}
}

func notImplementedParse(p *Parser, parent interface{}, token Token, literal string) error {
	pType := parent.(ElementI).GetElementType()
	return fmt.Errorf("parsing %s in %s not implemented yet", token.String(), pType.String())
}

func notValidChildParse(p *Parser, parent interface{}, token Token, literal string) error {
	pEType := parent.(ElementI).GetElementType().String()

	return FmtErrorf(p, "found %s, not a valid child for %s", token.String(), pEType)
}

func elementParse(p *Parser, parent interface{}, token Token, literal string) error {
	p.UnScan()
	var elem ElementI

	switch token {
	case ENTERPRISE:
		elem = NewEnterpriseStatement()
	case GROUP:
		elem = NewGroupStatement(parent.(ElementI).GetElementType())
	case PERSON:
		elem = NewPersonStatement()
	case SOFTWARE_SYSTEM:
		elem = NewSoftwareSystemStatement()
	case DEPLOYMENT_ENV:
		elem = NewDeploymentEnvironmentStatement()
	case ELEMENT:
		elem = NewElementStatement()
	}
	parent.(ElementContainer).AddElement(elem)
	return nextParse(elem.(Statement), p)
}

func closeBodyParse(p *Parser, parent interface{}, token Token, literal string) error {
	statement := parent.(Statement)
	if statement.BodyIsOpen() {
		statement.SetBodyOpen(false)
		return nil
	} else {
		return FmtErrorf(p, "unrecognized closing brace ('}'), not in statement body")
	}
}

func unexpectedEOFParse(p *Parser, parent interface{}, token Token, literal string) error {
	elem := parent.(ElementI)
	return FmtErrorf(p, "unexpected end of file in definition of %s", elem.GetElementType().String())
}
