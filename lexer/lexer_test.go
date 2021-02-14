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
		
	}
}