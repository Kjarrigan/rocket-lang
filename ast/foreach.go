package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Foreach struct {
	Token token.Token
	Index string
	Ident string
	Value Expression
	Body  *Block
}

func (fes *Foreach) expressionNode()      {}
func (fes *Foreach) TokenLiteral() string { return fes.Token.Literal }
func (fes *Foreach) String() string {
	var out bytes.Buffer
	out.WriteString("foreach ")
	out.WriteString(fes.Ident)
	out.WriteString(" ")
	out.WriteString(fes.Value.String())
	out.WriteString(fes.Body.String())
	return out.String()
}
