package lexer

import (
	"testing"
	"monkey/token"
)

func TestNextToken(t *testing){
	input := `=+(){},;`

	tests : = []struct {
		expectedType token.TokenType
		exectedLiteral string
	} {
		{ token.ASSIGN, "=" },
		{ token.PLUS, "+" },
		{ token.LPAREN, "(" },
		{ token.RPAREN, ")" },
		{ token.LBRACE, "{" },
		{ token.RBRACE, "}" },
		{ token.COMMA, "," },
		{ token.SEMICOLAN, ";"},
		{ token.EOF, ""}
		{}
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
	}
}