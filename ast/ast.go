// Package ast implements the Abstract Syntax Tree that represents the parsed
// source code before being passed on to the interpreter for evaluation.
package ast

import (
	"bytes"
	"github.com/ehsaaniqbal/ma/token"
)

// Defines an interface for all nodes in the AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Defines the interface for all statement nodes
type Statement interface {
	Node
	statementNode()
}

// Defines the interface for all expression nodes
type Expression interface {
	Node
	expressionNode()
}

// Root node of the AST
type Program struct {
	Statements []Statement
}

// Returns a stringified version of the AST for debugging
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Prints the literal value of the token associated with a specific node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
    Token token.Token
    ReturnValue Expression 
}

func (rs *ReturnStatement)  statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

