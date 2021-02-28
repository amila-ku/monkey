// Package ast(Abstract Syntax Tree)  parser starts here
package ast

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