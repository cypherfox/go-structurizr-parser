package parser

type Statement interface {
	Parse(p *Parser) error
}
