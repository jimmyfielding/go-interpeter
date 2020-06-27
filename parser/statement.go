package parser

import (
	"github.com/jimmyfielding/go-interpeter/ast"
	"github.com/jimmyfielding/go-interpeter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}
