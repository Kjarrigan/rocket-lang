package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseAssignExpression(name ast.Expression) ast.Expression {
	stmt := &ast.Assign{Token: p.curToken}
	if n, ok := name.(*ast.Identifier); ok {
		stmt.Name = n
	} else {
		msg := fmt.Sprintf("expected assign token to be IDENT, got %s instead", name.TokenLiteral())
		p.errors = append(p.errors, msg)
	}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}
