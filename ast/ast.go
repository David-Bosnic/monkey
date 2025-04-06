package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statement []Statement
}

type Identifier struct {
	Token token.Token //the token.IDENT token
	Value string      // Like the litteral string "taco" in taco = 5
}

type LetStatement struct { // Note the fact that we are kind pretending that they are all expressions for simplicity
	Token token.Token // The Let token
	Name  *Identifier // The Variable name IE like "x" in x = 5
	Value Expression  // expression is the generated value
}

func (ls *LetStatement) statementNode() {}

// Token Literal is the literal value of a token, so like dog in the string "dog"
// Token has also a field of type so it's type and literal
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string       { return "" }

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

func (l *Identifier) expressionNode()      {}
func (l *Identifier) TokenLiteral() string { return l.Token.Literal }
