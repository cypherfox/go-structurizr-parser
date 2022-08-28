package ast

type ElementType int32

const (
	SoftwareSystem ElementType = iota
)

type ElementContainer interface {
	GetElementByName(name string) Element
}

type Element interface {
	GetElementType() ElementType
	GetName() string
	AddTags(tags ...string)
}
