// Package token defines token struc tor lexer
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifier + Literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 1234

	//Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	LT		= "<"
	GT		= ">"


	//Delimiters
	COMMA     = ","
	SEMICOLAN = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	BANG	  = "!"
	ASTERISX	  = "*"
	SLASH		= "/"

	//keywords
	FUNCTION = "FUNCTION" // fn
	LET      = "LET"      // let
	IF		= "IF"	// if
	ELSE    = "ELSE" // else
	TRUE    = "TRUE" // true
	FALSE   = "FALSE" // false
	RETURN  = "RETURN" // return
)

// maps keyword string to tokey type definition
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

// LookupIdent returns tokentype for a given identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}
}
