package parser

import "github.com/jimmyfielding/go-interpeter/ast"

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}
