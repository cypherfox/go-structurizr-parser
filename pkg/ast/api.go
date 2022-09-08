package ast

import (
	. "github.com/cypherfox/go-structurizr-parser/pkg/parser"
)

type ElementType Token

const (
	Model          ElementType = ElementType(MODEL)
	SoftwareSystem ElementType = ElementType(SOFTWARE_SYSTEM)
	Enterprise     ElementType = ElementType(ENTERPRISE)
	Group          ElementType = ElementType(GROUP)
	Person         ElementType = ElementType(PERSON)
	Container      ElementType = ElementType(CONTAINER)
)

func (e ElementType) String() string {
	t := Token(e)
	return t.String()
}

type ElementContainer interface {
	GetElementByName(name string) Element
	AddElement(Element)
}

type Element interface {
	GetElementType() ElementType
	GetName() string
	AddTags(tags ...string) error
}
