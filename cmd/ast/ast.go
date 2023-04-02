package ast

import "github.com/EHughes190/monkey/cmd/token"

type Node interface {
	TokenLiteral() string
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
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the LET token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token // the RETURN token
	ReturnValue Expression
}

type Identifier struct {
	Token token.Token // the IDENT token
	Value string
}

// The program node is going to the root node of every AST
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// A basic let x = 5; could be represented by an AST of the form:
/*
                    _________________
                    |  *ast.Program |
                    |  Statements   |
                    _________________
                            |
                ____________|__________
                |  *ast.LetStatement |
                |         Name       |
                |        Value       |
                _____________________
                  |              |
                  |              |
  ________________|_            _|______________
  | *ast.Identifier|           |*ast.Expression|
  __________________           _________________

*/

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
