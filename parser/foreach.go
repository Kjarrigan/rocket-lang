package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseForEach() ast.Expression {
	expression := &ast.Foreach{Token: p.curToken}

	p.nextToken()
	expression.Ident = p.curToken.Literal

	if p.peekTokenIs(token.COMMA) {

		p.nextToken()

		if !p.peekTokenIs(token.IDENT) {
			p.errors = append(p.errors, fmt.Sprintf("second argument to foreach must be ident, got %v", p.peekToken))
			return nil
		}
		p.nextToken()

		expression.Index = expression.Ident
		expression.Ident = p.curToken.Literal

	}

	if !p.expectPeek(token.IN) {
		return nil
	}
	p.nextToken()

	expression.Value = p.parseExpression(LOWEST)
	if expression.Value == nil {
		return nil
	}

	p.nextToken()
	expression.Body = p.parseBlock()

	return expression
}