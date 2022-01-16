package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Assign struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (as *Assign) expressionNode()      {}
func (as *Assign) statementNode()       {}
func (as *Assign) TokenLiteral() string { return as.Token.Literal }
func (as *Assign) String() string {
	var out bytes.Buffer
	out.WriteString(as.Name.String())
	out.WriteString(" = ")
	out.WriteString(as.Value.String())
	return out.String()
}
