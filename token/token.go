package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//Identifier + literals
	IDENT = "IDENT" // add, foobar, x, y, etc.
	INT = "INT" // 123456

	//Operators
	ASSIGN = "="
	PLUS = "+"
	
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	// Keywords 
	FUCTION = "FUNCTION"
	LET = "LET"
)