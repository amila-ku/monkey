package parser

import (
	"github.com/amila-ku/monkey/lexer"
	"github.com/amila-ku/monkey/token"
	"github.com/amila-ku/monkey/ast"
)

type Parser struct {
	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read token twise to set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func(p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

//ParseProgram parses code and returns ast
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}